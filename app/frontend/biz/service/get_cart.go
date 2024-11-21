package service

import (
	"context"
	"strconv"

	common "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/YiD11/gomall/app/frontend/utils"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/cart"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var cartResp *cart.GetCartResp
	cartResp, err = rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context))})

	if err != nil {
		return nil, err
	}

	var items []map[string]string
	var total float64
	for _, it := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProducts(h.Context, &product.GetProductsReq{Id: it.ProductId})
		if err != nil {
			continue
		}
		var p *product.Product = productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Picture":     p.Picture,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Quantity":    strconv.Itoa(int(it.Quantity)),
		})
		total += float64(it.Quantity) * float64(p.Price)
	}
	return utils.H{
		"Title": "Cart",
		"Items": items,
		"Total": strconv.FormatFloat(total, 'f', 2, 64),
	}, nil
}
