package service

import (
	"context"
	"testing"

	"github.com/cloudwego/biz-demo/gomall/app/user/biz/dal/mysql"
	user "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		t.Logf("load env failed")
	}

	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "1demo@admin.com",
		Password:        "123456",
		PasswordConfirm: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}

// func TestRegisterService_Run(t *testing.T) {
// 	type args struct {
// 		req *user.RegisterReq
// 	}
// 	tests := []struct {
// 		name     string
// 		s        *RegisterService
// 		args     args
// 		wantResp *user.RegisterResp
// 		wantErr  bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotResp, err := tt.s.Run(tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("RegisterService.Run() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotResp, tt.wantResp) {
// 				t.Errorf("RegisterService.Run() = %v, want %v", gotResp, tt.wantResp)
// 			}
// 		})
// 	}
// }
