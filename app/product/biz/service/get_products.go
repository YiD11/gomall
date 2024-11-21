package service

import (
	"context"

	"github.com/YiD11/gomall/app/product/biz/dal/mysql"
	"github.com/YiD11/gomall/app/product/biz/dal/redis"
	"github.com/YiD11/gomall/app/product/biz/model"
	product "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductsService struct {
	ctx context.Context
} // NewGetProductsService new GetProductsService
func NewGetProductsService(ctx context.Context) *GetProductsService {
	return &GetProductsService{ctx: ctx}
}

// Run create note info
func (s *GetProductsService) Run(req *product.GetProductsReq) (resp *product.GetProductsResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}

	productQuery := model.NewCachedProductQuery(s.ctx, mysql.DB, redis.RedisClient)

	p, err := productQuery.GetById(int(req.Id))
	if err != nil {
		return nil, err
	}

	return &product.GetProductsResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Price:       p.Price,
			Picture:     p.Picture,
			Description: p.Description,
			Name:        p.Name,
		},
	}, nil
}
