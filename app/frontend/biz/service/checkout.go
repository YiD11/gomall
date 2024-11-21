package service

import (
	"context"
	"strconv"

	common "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/YiD11/gomall/app/frontend/utils"
	rpccart "github.com/YiD11/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var items []map[string]string
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	CartResp, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var total float32

	for _, it := range CartResp.Items{
		productResp, err := rpc.ProductClient.GetProducts(h.Context, &rpcproduct.GetProductsReq{Id: it.ProductId})
		if err != nil {
			return nil, err
		}
		p := productResp.Product
		if p == nil {
			continue
		}
		items = append(items, map[string]string{
			"Name": p.Name,
			"Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Quantity": strconv.Itoa(int(it.Quantity)),
		})
		total += float32(it.Quantity) * p.Price
	}

	return utils.H{
		"Title": "Checkout",
		"Items": items,
		"Total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
