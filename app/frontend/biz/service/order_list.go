package service

import (
	"context"
	"time"

	common "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	"github.com/YiD11/gomall/app/frontend/types"
	frontendUtils "github.com/YiD11/gomall/app/frontend/utils"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/order"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	userId := frontendUtils.GetUserIdFromCtx(h.Context)

	var p *order.ListOrdersResp

	p, err = rpc.OrderClient.ListOrders(h.Context, &order.ListOrdersReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var list []*types.Order
	for _, v := range p.Orders {
		var (
			total float32
			items []types.OrderItem
		)

		for _, it := range v.Items {
			total += it.Cost * float32(it.Item.Quantity)

			prodResp, err := rpc.ProductClient.GetProducts(h.Context, &product.GetProductsReq{Id: it.Item.ProductId})
			if err != nil {
				return nil, err
			}

			if prodResp == nil || prodResp.Product == nil {
				continue
			}

			items = append(items, types.OrderItem{
				ProductName: prodResp.Product.Name,
				Picture:     prodResp.Product.Picture,
				Quantity:    it.Item.Quantity,
				Cost:        it.Cost,
			})
		}

		createdTime := time.Unix(v.CreatedAt, 0)

		list = append(list, &types.Order{
			OrderId:     v.OrderId,
			CreatedDate: createdTime.Format("2006-01-02 15:04:05"),
			Cost:        total,
			Items:       items,
		})
	}

	return utils.H{
		"Title":  "Order",
		"Orders": list,
	}, nil
}
