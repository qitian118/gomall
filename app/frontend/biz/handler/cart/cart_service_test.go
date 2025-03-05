package cart

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestAddCartItem(t *testing.T) {
	h := server.Default()
	h.POST("/cart", AddCartItem)
	path := "/cart" // todo: you can customize query
	// 模拟添加购物车的请求体，包含 user_id 和 CartItem（product_id 和 quantity）
	body := &ut.Body{
		Body: bytes.NewBufferString(`{
			"user_id": 1, 
			"item": {
				"product_id": 123, 
				"quantity": 2
			}
		}`),
		Len: len(`{
			"user_id": 1, 
			"item": {
				"product_id": 123, 
				"quantity": 2
			}
		}`),
	}

	header := ut.Header{} // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	assert.DeepEqual(t, 200, resp.StatusCode())

}

func TestGetCart(t *testing.T) {
	h := server.Default()
	h.GET("/cart", GetCart)
	path := "/cart" // todo: you can customize query
	body := &ut.Body{
		Body: bytes.NewBufferString(`{
			"user_id": 1
		}`),
		Len: len(`{
			"user_id": 1
		}`),
	}

	header := ut.Header{} // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, true, bytes.Contains(resp.Body(), []byte("items")))
	assert.DeepEqual(t, true, bytes.Contains(resp.Body(), []byte("total")))
}
