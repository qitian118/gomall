package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/qitian118/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/qitian118/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/qitian118/gomall/app/frontend/utils"
	"github.com/qitian118/gomall/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEmptyCartService(Context context.Context, RequestContext *app.RequestContext) *EmptyCartService {
	return &EmptyCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EmptyCartService) Run(req *common.Empty) (resp bool, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	_, err = rpc.CartClient.EmptyCart(h.Context, &cart.EmptyCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return false, err
	}

	return true, nil
}
