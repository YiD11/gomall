package order

import (
	"context"
	order "github.com/YiD11/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderReq, callOptions ...callopt.Option) (resp *order.PlaceOrderResp, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOrders(ctx context.Context, req *order.ListOrdersReq, callOptions ...callopt.Option) (resp *order.ListOrdersResp, err error) {
	resp, err = defaultClient.ListOrders(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOrders call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
