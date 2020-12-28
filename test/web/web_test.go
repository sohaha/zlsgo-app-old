package web_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
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

func walkDir(f func(path string)) {
	_ = filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		f(path)
		return nil
	})
}

func TestMain(m *testing.M) {
	fileList := make([]string, 0)

	walkDir(func(path string) {
		fileList = append(fileList, path)
	})

	// 如果项目根目录存在配置文件，直接 copy 一份
	if zfile.FileExist("../../conf.yml") {
		_ = zfile.CopyFile("../../conf.yml", "./conf.yml")
	}

	// 清除数据
	defer func() {
		walkDir(func(path string) {
			rm := true
			for i, v := range fileList {
				if path == v {
					rm = false
					fileList = append(fileList[:i], fileList[i+1:]...)
					break
				}
			}
			if rm {
				zfile.Rmdir(path)
			}
		})
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
