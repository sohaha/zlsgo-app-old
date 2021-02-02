package web_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/global/task"
	"app/test"
	"app/web/router"
	. "github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
)

type stBody struct {
	Body        io.Reader
	ContentType string
	Header      map[string]string
}

var net *znet.Engine

func request(method, url string, body *stBody) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body.Body)
	if body.Header != nil {
		for k, v := range body.Header {
			req.Header.Set(k, v)
		}
	}
	net.ServeHTTP(w, req)
	return w
}

func TestMain(m *testing.M) {
	test.Run(func() {
		// 初始化定时任务
		task.Init()

		// 初始化 Web 服务
		router.Init()

		// 获取服务
		net = router.Engine

		// 运行测试
		m.Run()
	})
}

func TestHome(t *testing.T) {
	tt := NewTest(t)

	w := request("GET", "/", &stBody{})

	res := w.Body.String()
	tt.EqualExit(200, w.Code)
	tt.EqualExit(200, zjson.Get(res, "code").Int())

	t.Log("测试通过", res)
}
