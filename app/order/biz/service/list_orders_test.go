package service

import (
	"context"
	"testing"
	order "github.com/YiD11/gomall/rpc_gen/kitex_gen/order"
)

func TestListOrders_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOrdersService(ctx)
	// init req and assert value

	req := &order.ListOrdersReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
