package service

import (
	"context"
	"time"

	auth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
	"github.com/golang-jwt/jwt/v5"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

type CustomClaims struct {
	UserId   int32  `json:"user_id"`
	UserRole string `json:"user_role"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
	jwt.RegisteredClaims
}

var secretKey = []byte("your-256-bit-secret")

// Run create note info
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	// Finish your business logic.
	//这里中间件会放行分发操作

	claims := CustomClaims{
		UserId: req.UserId,
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
