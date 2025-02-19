package service

import (
	"context"

	home "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	resp := make(map[string]any)

	items := []map[string]any{
		{"Name": "t-shirt1", "Price": 100, "Picture": "/static/image/t-shirt1.png"},
		{"Name": "t-shirt2", "Price": 101, "Picture": "/static/image/t-shirt2.png"},
		{"Name": "t-shirt3", "Price": 102, "Picture": "/static/image/t-shirt3.png"},
		{"Name": "t-shirt4", "Price": 103, "Picture": "/static/image/t-shirt4.png"},
		{"Name": "t-shirt5", "Price": 104, "Picture": "/static/image/t-shirt5.png"},
		{"Name": "t-shirt6", "Price": 105, "Picture": "/static/image/t-shirt6.png"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
