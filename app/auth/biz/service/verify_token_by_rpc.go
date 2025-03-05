package service

import (
	"context"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	auth "github.com/qitian118/gomall/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// Finish your business logic.

	tokenSigned := req.Token
	if tokenSigned == "" {
		return &auth.VerifyResp{Res: false}, errors.New("token isn't exist")
	}

	token, err := jwt.Parse(tokenSigned, func(t *jwt.Token) (interface{}, error) {
		return []byte("ISHDSDAJKDH"), nil
	})

	if err != nil || !token.Valid {
		return &auth.VerifyResp{Res: false}, errors.New("invalid token")
	}

	return &auth.VerifyResp{Res: true}, nil
}
