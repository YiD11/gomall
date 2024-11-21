package home

import (
	"context"

	"github.com/YiD11/gomall/app/frontend/biz/service"
	"github.com/YiD11/gomall/app/frontend/biz/utils"
	common "github.com/YiD11/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Home .
// @router / [GET]
func Home(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHomeService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "home.tmpl", utils.WarpResponse(ctx, c, resp))

	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
