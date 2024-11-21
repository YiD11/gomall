package service

import (
	"context"

	"github.com/YiD11/gomall/app/cart/biz/dal/mysql"
	"github.com/YiD11/gomall/app/cart/biz/model"
	cart "github.com/YiD11/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	rows, err := model.GetCartById(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50002, err.Error())
	}
	var items []*cart.CartItem
	for _, it := range rows {
		items = append(items, &cart.CartItem{
			ProductId: it.ProductId,
			Quantity: it.Quantity,
		})
	}
	return &cart.GetCartResp{Items: items}, nil
}
