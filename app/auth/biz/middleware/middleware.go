package middleware

import (
	"context"
	"errors"

	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuthMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		// 判断是否为分发token请求
		if _, isDeliver := req.(*auth.DeliverTokenReq); isDeliver {
			// 直接放行分发token请求
			return next(ctx, req, resp)
		}

		// 获取 JWT Token
		authReq := req.(*auth.VerifyTokenReq)
		tokenStr := authReq.Token
		if tokenStr == "" {
			return errors.New("token required")
		}

		// 解析并验证 Token
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			// todo 暂时写死
			return []byte("your-256-bit-secret"), nil
		})
		if err != nil || !token.Valid {
			return errors.New("invalid token")
		}

		// 将 claims 存入上下文
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			ctx = context.WithValue(ctx, "user_role", claims["user_role"])
			ctx = context.WithValue(ctx, "resource", claims["resource"])
			ctx = context.WithValue(ctx, "action", claims["action"])

		}

		return next(ctx, req, resp)
	}
}

func CasbinMiddleware(enforcer *casbin.Enforcer) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			// 获取上下文中的角色、资源、动作
			userRole := ctx.Value("user_role").(string)
			resource := ctx.Value("resource").(string)
			action := ctx.Value("action").(string)

			// 判断是否有权限
			ok, err := enforcer.Enforce(userRole, resource, action)
			if err != nil || !ok {
				return errors.New("forbidden")
			}

			return next(ctx, req, resp)
		}
	}
}

func ChainMiddleware(middlewares ...endpoint.Middleware) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
