package utils

import (
	"context"

	rpcauth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
)

func GetToken(userid int32, c context.Context) (string, error) {
	deliveryResp, err := rpc.AuthClient.DeliverTokenByRPC(c, &rpcauth.DeliverTokenReq{UserId: userid})
	if err != nil {
		return "", err
	}
	return deliveryResp.Token, nil
}
