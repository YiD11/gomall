package service

import (
	"context"

	category "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	var p *rpcproduct.ListProductsResp
	p, err = rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
