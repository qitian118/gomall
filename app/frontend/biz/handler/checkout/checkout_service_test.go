package checkout

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestCheckout(t *testing.T) {
	h := server.Default()
	h.GET("/checkout", Checkout)
	path := "/checkout"                                       // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, true, bytes.Contains(resp.Body(), []byte("Checkout")))
}

func TestCheckoutWaiting(t *testing.T) {
	h := server.Default()
	h.POST("/checkout/waiting", CheckoutWaiting)
	path := "/checkout/waiting" // todo: you can customize query
	body := &ut.Body{
		Body: bytes.NewBufferString(`{
			"email": "testuser@example.com",
			"firstname": "John",
			"lastname": "Doe",
			"street": "123 Main St",
			"zipcode": "12345",
			"province": "SomeProvince",
			"country": "SomeCountry",
			"city": "SomeCity",
			"card_num": "1234567812345678",
			"expiration_month": 12,
			"expiration_year": 2025,
			"cvv": 123,
			"payment": "credit_card"
		}`),
		Len: len(`{
			"email": "testuser@example.com",
			"firstname": "John",
			"lastname": "Doe",
			"street": "123 Main St",
			"zipcode": "12345",
			"province": "SomeProvince",
			"country": "SomeCountry",
			"city": "SomeCity",
			"card_num": "1234567812345678",
			"expiration_month": 12,
			"expiration_year": 2025,
			"cvv": 123,
			"payment": "credit_card"
		}`),
	}
	header := ut.Header{} // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "POST", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	// todo edit your unit test.
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, true, bytes.Contains(resp.Body(), []byte("waiting")))
}

func TestCheckoutResult(t *testing.T) {
	h := server.Default()
	h.GET("/checkout/result", CheckoutResult)
	path := "/checkout/result"                                // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))
	assert.DeepEqual(t, 200, resp.StatusCode())

}
