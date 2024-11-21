package main

import (
	"context"
	order "github.com/YiD11/gomall/rpc_gen/kitex_gen/order"
	"github.com/YiD11/gomall/app/order/biz/service"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// PlaceOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	resp, err = service.NewPlaceOrderService(ctx).Run(req)

	return resp, err
}

// ListOrders implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrders(ctx context.Context, req *order.ListOrdersReq) (resp *order.ListOrdersResp, err error) {
	resp, err = service.NewListOrdersService(ctx).Run(req)

	return resp, err
}
