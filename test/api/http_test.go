package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"app/service"
	"app/service/router"
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
	// 初始化
	service.InitEngine()
	// 获取服务
	net = router.Engine
	// 运行测试
	s := m.Run()
	// 清除数据
	service.Clear()
	os.Exit(s)
}

func TestHome(t *testing.T) {
	tt := NewTest(t)

	w := request("GET", "/", nil)

	res := w.Body.String()
	tt.EqualExit(200, w.Code)
	tt.EqualExit(200, zjson.Get(res, "code").Int())

	t.Log("测试通过", res)
}
