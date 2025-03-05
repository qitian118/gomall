package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/user/biz/model"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
)

type UserInfoService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

// Run create note info
func (s *UserInfoService) Run(req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	// Finish your business logic.

	id := req.UserId

	row, err := model.GetByUserid(mysql.DB, uint(id))
	if err != nil {
		return nil, err
	}

	resp = &user.UserInfoResp{
		Email: row.Email,
	}
	// if req.Email == "" || req.Password == "" {
	// 	return nil, errors.New("email or password is empty")
	// }
	// row, err := model.GetByEmail(mysql.DB, req.Email)
	// if err != nil {
	// 	return nil, err
	// }
	// err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	// if err != nil {
	// 	return nil, err
	// }
	// resp = &user.LoginResp{
	// 	UserId: int32(row.ID),
	// }
	return resp, nil
}
