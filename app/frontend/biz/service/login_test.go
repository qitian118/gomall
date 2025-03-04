package service

import (
	"context"
	"testing"

	auth "github.com/gomall_project/app/frontend/hertz_gen/frontend/auth"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginService(ctx, nil)
	// init req and assert value

	req := &auth.LoginReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
