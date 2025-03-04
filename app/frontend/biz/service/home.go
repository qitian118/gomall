package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/gomall_project/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "T-shirt1", "Price": 100, "Picture": "/static/image/movie.jpg"},
		{"Name": "T-shirt2", "Price": 120, "Picture": "/static/image/yundong.jpg"},
		{"Name": "T-shirt3", "Price": 140, "Picture": "/static/image/yundong.jpg"},
		{"Name": "T-shirt4", "Price": 150, "Picture": "/static/image/movie.jpg"},
		{"Name": "T-shirt5", "Price": 200, "Picture": "/static/image/movie.jpg"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items
	return resp, nil
}
