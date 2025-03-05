package home

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	//"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestHome(t *testing.T) {
	h := server.Default()
	h.GET("/", Home)
	path := "/"                                               // todo: you can customize query
	body := &ut.Body{Body: bytes.NewBufferString(""), Len: 1} // todo: you can customize body
	header := ut.Header{}                                     // todo: you can customize header
	w := ut.PerformRequest(h.Engine, "GET", path, body, header)
	resp := w.Result()
	t.Log(string(resp.Body()))

	assert.DeepEqual(t, 200, resp.StatusCode())

	// 检查响应体中是否包含 title
	assert.DeepEqual(t, true, bytes.Contains(resp.Body(), []byte("Hot sale")))

	assert.DeepEqual(t, true, bytes.Contains(resp.Body(), []byte("items")))
}
