package service

import (
	"context"

	"github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/cart"
	"github.com/YiD11/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/YiD11/gomall/app/frontend/utils"
	rpccart "github.com/YiD11/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// log.Println(req.ProductNum, req.ProductId)
	if req.ProductNum <= 0 || req.ProductId == 0 {
		return nil, kerrors.NewBizStatusError(50003, "request unvalid")
	}
	_, err = rpc.CartClient.AddItem(
		h.Context,
		&rpccart.AddItemReq{
			UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
			Item:   &rpccart.CartItem{ProductId: req.ProductId, Quantity: uint32(req.ProductNum)},
		},
	)
	if err != nil {
		return nil, err
	}
	return
}
