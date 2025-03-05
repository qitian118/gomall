package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/qitian118/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/qitian118/gomall/app/frontend/utils"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/auth"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")
		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从Header获取token
		s := sessions.Default(c)
		// s.Set("next", c.Request.URI().String())
		userIDVal := s.Get("user_id")
		if userIDVal == nil {

			// c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": "Missing user ID"})
			// fmt.Println("after abort")
			c.Redirect(302, []byte("/sign-in"))
			c.Abort()
			// fmt.Println("after redirect")

			return
		}
		userId, ok := userIDVal.(int32)
		if !ok {
			c.Redirect(302, []byte("/sign-in"))
			c.Abort()
			return
		}
		tokenVal := s.Get("token")
		if tokenVal == nil {
			c.Redirect(302, []byte("/sign-in"))
			c.Abort()
			return
		}
		token, ok := tokenVal.(string)
		if !ok {
			c.Redirect(302, []byte("/sign-in"))
			c.Abort()
			return
		}

		// 调用Auth服务验证
		resp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{
			Token: token,
		})

		if err != nil || !resp.Res {
			c.Redirect(302, []byte("/sign-in"))
			c.Abort()
			return
		}

		// 注入用户信息到上下文
		c.Set("user_id", userId)
		c.Next(ctx)
	}
}
