package home

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/gomall_project/app/frontend/biz/service"
	"github.com/gomall_project/app/frontend/biz/utils"
	common "github.com/gomall_project/app/frontend/hertz_gen/frontend/common"
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
	// 返回html页面
	// utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
	resp["user_id"] = 222
	// 加这行日志,这是用来打印resp的
	c.HTML(consts.StatusOK, "home.html", utils.WrapResponse(ctx, c, resp)) // 返回html页面, consts.StatusOK 状态码,用于判断是否成功， "home.tmpl" 模板名称, resp 数据，用于渲染模板
}
