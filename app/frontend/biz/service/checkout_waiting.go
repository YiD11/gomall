package service

import (
	"context"

	checkout "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/YiD11/gomall/app/frontend/utils"
	rpccheckout "github.com/YiD11/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	_, err = rpc.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId:    uint32(userId),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Address: &rpccheckout.Address{
			Country:       req.Country,
			State:         req.Province,
			City:          req.City,
			ZipCode:       req.Zipcode,
			StreetAddress: req.Street,
		},
		CreditCard: &payment.CreditCardInfo{
			Number:          req.CardNum,
			Cvv:             req.Cvv,
			ExpirationYear:  req.ExpirationYear,
			ExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"Title":    "Waiting",
		"Redirect": "/checkout/result",
	}, nil
}
