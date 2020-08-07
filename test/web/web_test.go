package web_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/module/initialize"
	"app/web/router"
	. "github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
)

var net *znet.Engine

func request(method, url string, body io.Reader) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	net.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	// 清除数据
	defer initialize.Clear()
	// 初始化
	initialize.InitEngine()
	// 获取服务
	net = router.Engine
	// 运行测试
	m.Run()
}

func TestHome(t *testing.T) {
	tt := NewTest(t)

	w := request("GET", "/", nil)

	res := w.Body.String()
	tt.EqualExit(200, w.Code)
	tt.EqualExit(200, zjson.Get(res, "code").Int())

	t.Log("测试通过", res)
}
