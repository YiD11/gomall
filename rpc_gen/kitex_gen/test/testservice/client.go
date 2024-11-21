// Code generated by Kitex v0.9.1. DO NOT EDIT.

package testservice

import (
	"context"
	test "github.com/YiD11/gomall/rpc_gen/kitex_gen/test"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PrintTest(ctx context.Context, Req *test.TestReq, callOptions ...callopt.Option) (r *test.TestResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kTestServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kTestServiceClient struct {
	*kClient
}

func (p *kTestServiceClient) PrintTest(ctx context.Context, Req *test.TestReq, callOptions ...callopt.Option) (r *test.TestResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PrintTest(ctx, Req)
}
