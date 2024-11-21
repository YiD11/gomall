package service

import (
	"context"
	"testing"
	product "github.com/YiD11/gomall/rpc_gen/kitex_gen/product"
)

func TestGetProducts_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductsService(ctx)
	// init req and assert value

	req := &product.GetProductsReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
