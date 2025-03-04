package service

import (
	"context"
	"testing"

	"github.com/gomall_project/app/user/biz/dal/mysql"
	user "github.com/gomall_project/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatalf("加载环境变量失败: %v", err)
	}
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:    "test@test.com",
		Password: "123456",
		PasswordConfirm: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
