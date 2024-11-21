package service

import (
	"context"

	"github.com/YiD11/gomall/app/cart/biz/dal/mysql"
	"github.com/YiD11/gomall/app/cart/biz/model"
	"github.com/YiD11/gomall/app/cart/infra/rpc"
	cart "github.com/YiD11/gomall/rpc_gen/kitex_gen/cart"
	"github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	var p *product.GetProductsResp
	p, err = rpc.ProductClient.GetProducts(s.ctx, &product.GetProductsReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if p == nil || p.Product.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(40004, "product not found")
	}

	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Quantity:  req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	
	return &cart.AddItemResp{}, nil
}
