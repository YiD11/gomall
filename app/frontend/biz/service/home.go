package service

import (
	"context"

	common "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var p *rpcproduct.ListProductsResp
	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{CategoryName: "T-shirt"})
	if err != nil {
		return nil, err
	}

	resp := utils.H{
		"Title": "YiD11",
		"Items": p.Products,
	}

	return resp, nil
}
