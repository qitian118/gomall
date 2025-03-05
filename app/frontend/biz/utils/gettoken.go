package utils

import (
	"context"

	rpcauth "github.com/qitian118/gomall/rpc_gen/kitex_gen/auth"

	"github.com/qitian118/gomall/app/frontend/infra/rpc"
)

func GetToken(userid int32, c context.Context) (string, error) {
	deliveryResp, err := rpc.AuthClient.DeliverTokenByRPC(c, &rpcauth.DeliverTokenReq{UserId: userid})
	if err != nil {
		return "", err
	}
	return deliveryResp.Token, nil
}
