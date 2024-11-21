package test

import (
	"context"
	test "github.com/YiD11/gomall/rpc_gen/kitex_gen/test"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PrintTest(ctx context.Context, req *test.TestReq, callOptions ...callopt.Option) (resp *test.TestResp, err error) {
	resp, err = defaultClient.PrintTest(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PrintTest call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
