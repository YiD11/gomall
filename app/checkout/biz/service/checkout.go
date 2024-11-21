package service

import (
	"context"
	"log"

	"github.com/YiD11/gomall/app/checkout/infra/mq"
	"github.com/YiD11/gomall/app/checkout/infra/rpc"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/YiD11/gomall/rpc_gen/kitex_gen/checkout"
	email "github.com/YiD11/gomall/rpc_gen/kitex_gen/email"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/order"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/payment"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	cartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResp == nil || cartResp.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	var (
		total float32
		oi    []*order.OrderItem
	)

	for _, it := range cartResp.Items {
		pid := it.ProductId

		productResp, err := rpc.ProductClient.GetProducts(s.ctx, &product.GetProductsReq{Id: pid})
		if err != nil {
			return nil, err
		}
		if productResp == nil || productResp.Product == nil {
			continue
		}

		price := productResp.Product.Price
		cost := float32(it.Quantity) * price
		total += cost

		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{
				ProductId: pid,
				Quantity:  it.Quantity,
			},
			Cost: cost,
		})
	}

	var p *order.PlaceOrderResp
	p, err = rpc.OrderClient.PlaceOrder(s.ctx, &order.PlaceOrderReq{
		UserId: req.UserId,
		Email:  req.Email,
		Address: &order.Address{
			StreetAddress: req.Address.StreetAddress,
			City:          req.Address.City,
			State:         req.Address.State,
			Country:       req.Address.Country,
			ZipCode:       req.Address.ZipCode,
		},
		Items: oi,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5004002, err.Error())
	}

	var orderId string
	if p.Order != nil && p != nil {
		orderId = p.Order.OrderId
	}

	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			Number:          req.CreditCard.Number,
			Cvv:             req.CreditCard.Cvv,
			ExpirationYear:  req.CreditCard.ExpirationYear,
			ExpirationMonth: req.CreditCard.ExpirationMonth,
		},
	}

	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	payResp, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	klog.Info(payResp)

	emailReq := &email.EmailReq{
		From:        "noreply@gomall.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "An order created by you in gomall",
		Content:     "An order created by you in gomall",
	}

	data, _ := proto.Marshal(emailReq)

	msg := &nats.Msg{
		Subject: "email",
		Data:    data,
		Header:  make(nats.Header),
	}
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	err = mq.Nc.PublishMsg(msg)
	if err != nil {
		log.Println("Failed to publish message:", err.Error())
	}

	return &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: payResp.TransactionId,
	}, nil
}
