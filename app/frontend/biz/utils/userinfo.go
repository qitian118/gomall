package utils

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/qitian118/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/qitian118/gomall/app/frontend/utils"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/user"
)

func UserInfo(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	userId := frontendUtils.GetUserIdFromCtx(ctx)
	content["user_id"] = userId
	// fmt.Println("berfore userclient")
	resp, err := rpc.UserClient.UserInfo(ctx, &user.UserInfoReq{UserId: userId})
	// fmt.Println("after userclient")
	if err != nil {
		fmt.Println("have error")
		fmt.Println(err)
		return content
	}
	content["email"] = resp.Email
	return content
}
