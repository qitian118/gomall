package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/qitian118/gomall/app/frontend/biz/utils"
	auth "github.com/qitian118/gomall/app/frontend/hertz_gen/frontend/auth"
	"github.com/qitian118/gomall/app/frontend/infra/rpc"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/user"
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
	token, err := utils.GetToken(resp.UserId, h.Context)
	if err != nil {
		return "deliver token failed", err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	session.Set("token", token)

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
