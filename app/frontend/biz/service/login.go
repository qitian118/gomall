package service

import (
	"context"
	"fmt"

	auth "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	rpcauth "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo  user sev api
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "login failed", err
	}
	fmt.Println("a new login")
	deliveryResp, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, &rpcauth.DeliverTokenReq{UserId: resp.UserId})
	fmt.Println("after deliver")
	if err != nil {
		fmt.Println(err)
		return "deliver token failed", err
	}
	fmt.Println("get a token")
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	session.Set("token", deliveryResp.Token)
	fmt.Println(deliveryResp.Token)
	err = session.Save()
	if err != nil {
		return "", err
	}
	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
