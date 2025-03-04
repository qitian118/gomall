package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default()

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.Data(consts.StatusOK, consts.MIMETextPlain, []byte("hello world")) //

	})

	h.Spin()
}
