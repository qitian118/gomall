package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"github.com/qitian118/gomall/app/frontend/biz/utils"
	auth "github.com/qitian118/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/qitian118/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/qitian118/gomall/app/frontend/infra/rpc"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return nil, err
	}
	// fmt.Println("a new register")
	token, err := utils.GetToken(userResp.UserId, h.Context)
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	session.Set("token", token)
	err = session.Save()
	if err != nil {
		return nil, err
	}

	return
}
