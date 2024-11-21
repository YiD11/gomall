package test

import (
	"context"
	test "github.com/YiD11/gomall/rpc_gen/kitex_gen/test"

	"github.com/YiD11/gomall/rpc_gen/kitex_gen/test/testservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() testservice.Client
	Service() string
	PrintTest(ctx context.Context, Req *test.TestReq, callOptions ...callopt.Option) (r *test.TestResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := testservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient testservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() testservice.Client {
	return c.kitexClient
}

func (c *clientImpl) PrintTest(ctx context.Context, Req *test.TestReq, callOptions ...callopt.Option) (r *test.TestResp, err error) {
	return c.kitexClient.PrintTest(ctx, Req, callOptions...)
}
