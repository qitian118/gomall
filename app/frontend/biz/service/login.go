package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/gomall_project/app/frontend/hertz_gen/frontend/auth"
	"github.com/gomall_project/app/frontend/infra/rpc"

	"github.com/hertz-contrib/sessions"

	"github.com/gomall_project/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

// Run create note info
func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	// Finish your business logic.

	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return "", err
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
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
