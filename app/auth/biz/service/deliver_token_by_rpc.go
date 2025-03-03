package service

import (
	"context"
	"time"

	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v5"
)

// 自定义声明结构
type CustomClaims struct {
	UserId   int32  `json:"user_id"`
	UserRole string `json:"user_role"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
	jwt.RegisteredClaims
}

// todo 暂时写死
var secretKey = []byte("your-256-bit-secret")

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	// 这里中间件会放行分发操作

	claims := CustomClaims{
		UserId:   req.UserId,
		UserRole: req.UserRole,
		Resource: req.Resource,
		Action:   req.Action,

		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "your-app-name",
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return nil, err
	}
	resp = &auth.DeliveryResp{
		Token: token,
	}

	return resp, nil
}
