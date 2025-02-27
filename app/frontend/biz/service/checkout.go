package service

import (
	"context"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/infra/rpc"
	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/checkout"
	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	rpccart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var items []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: userId})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
