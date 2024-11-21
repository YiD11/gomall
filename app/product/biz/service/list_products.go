package service

import (
	"context"

	"github.com/YiD11/gomall/app/product/biz/dal/mysql"
	"github.com/YiD11/gomall/app/product/biz/model"
	product "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// if req.CategoryName == "" {
	// 	return nil, kerrors.NewGRPCBizStatusError(2004001, "category name is required")
	// }

	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCateGoryName(req.CategoryName)
	// if err != nil {
	// 	return nil, err
	// }

	resp = &product.ListProductsResp{}
	for _, vl := range c {
		for _, p := range vl.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Price:       p.Price,
				Picture:     p.Picture,
				Description: p.Description,
				Name:        p.Name,
			})
		}
	}
	return resp, nil
}
