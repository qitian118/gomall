package service

import (
	"context"
	"errors"

	"github.com/gomall_project/app/user/biz/dal/mysql"
	"github.com/gomall_project/app/user/biz/model"
	user "github.com/gomall_project/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("邮箱、密码不能为空")
	}

	row, err := model.GetByEmail(mysql.DB, req.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	resp = &user.LoginResp{
		UserId: int32(row.ID),
	}

	return resp, nil
}
