package service

import (
	"context"

	"github.com/YiD11/gomall/app/product/biz/dal/mysql"
	"github.com/YiD11/gomall/app/product/biz/model"
	product "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	productQuery := model.NewProductQuery(s.ctx, mysql.DB)

	ps, err := productQuery.SearchProducts(req.Query)

	resp = &product.SearchProductsResp{}
	for _, p := range ps {
		resp.Products = append(resp.Products, &product.Product{
			Id:          uint32(p.ID),
			Price:       p.Price,
			Picture:     p.Picture,
			Description: p.Description,
			Name:        p.Name,
		})
	}
	return
}
