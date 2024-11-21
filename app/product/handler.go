package main

import (
	"context"
	product "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
	"github.com/YiD11/gomall/app/product/biz/service"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProducts(ctx context.Context, req *product.GetProductsReq) (resp *product.GetProductsResp, err error) {
	resp, err = service.NewGetProductsService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}
