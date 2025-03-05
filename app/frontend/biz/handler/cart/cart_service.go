package cart

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/qitian118/gomall/app/frontend/biz/service"
	"github.com/qitian118/gomall/app/frontend/biz/utils"
	cart "github.com/qitian118/gomall/app/frontend/hertz_gen/frontend/cart"
	common "github.com/qitian118/gomall/app/frontend/hertz_gen/frontend/common"
)

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "cart", utils.WarpResponse(ctx, c, resp))
}

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.Redirect(consts.StatusFound, []byte("/cart"))
}

func EmptyCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	utils.SendErrResponse(ctx, c, consts.StatusOK, err)
	// 	return
	// }

	yn, err := service.NewEmptyCartService(ctx, c).Run(&req)
	if err != nil || !yn {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	// 给客户端一个成功的提示
	c.Set("success_message", "商品已成功添加到购物车")
	c.Redirect(consts.StatusFound, []byte("/cart"))
}
