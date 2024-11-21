package service

import (
	"context"

	"github.com/YiD11/gomall/app/order/biz/dal/mysql"
	"github.com/YiD11/gomall/app/order/biz/model"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/YiD11/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrdersService struct {
	ctx context.Context
} // NewListOrdersService new ListOrdersService
func NewListOrdersService(ctx context.Context) *ListOrdersService {
	return &ListOrdersService{ctx: ctx}
}

// Run create note info
func (s *ListOrdersService) Run(req *order.ListOrdersReq) (resp *order.ListOrdersResp, err error) {
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500002, err.Error())
	}

	var orders []*order.Order
	for _, v := range list {
		var items []*order.OrderItem
		for _, oi := range v.OrderItems {
			items = append(items, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity: oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}

		orders = append(orders, &order.Order{
			OrderId: v.OrderId,
			UserId: v.UserId,
			CreatedAt: v.CreatedAt.Unix(),
			UserCurrency: v.UserCurrency,
			Email: v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City: v.Consignee.City,
				State: v.Consignee.State,
				Country: v.Consignee.Country,
				ZipCode: v.Consignee.ZipCode,
			},
			Items: items,
		})
	}
	resp = &order.ListOrdersResp{
		Orders: orders,
	}
	return
}
