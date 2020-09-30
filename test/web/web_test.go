package web_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"

	"app/global/initialize"
	"app/web/router"
)

var net *znet.Engine

type stBody struct {
	Body        io.Reader
	ContentType string
	Header      map[string]string
}

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
	// 清除数据
	defer func() {
		zfile.Rmdir("./db.sqlite")
		// initialize.Clear()
	}()

	// 初始化
	initialize.InitEngine()
	// 获取服务
	net = router.Engine
	// 运行测试
	m.Run()
}

func TestHome(t *testing.T) {
	tt := NewTest(t)

	w := request("GET", "/", &stBody{})

	res := w.Body.String()
	tt.EqualExit(200, w.Code)
	tt.EqualExit(200, zjson.Get(res, "code").Int())

	t.Log("测试通过", res)
}
