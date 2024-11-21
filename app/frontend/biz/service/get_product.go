package service

import (
	"context"

	product "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	var p *rpcproduct.GetProductsResp
	p, err = rpc.ProductClient.GetProducts(h.Context, &rpcproduct.GetProductsReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"item": p.Product,
	}, nil
}
