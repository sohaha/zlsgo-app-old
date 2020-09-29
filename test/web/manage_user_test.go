package web_test

import (
	"strings"
	"testing"

	. "github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zjson"
)

func TestUserLogin(t *testing.T) {
	tt := NewTest(t)
	body := strings.NewReader("user=manage&pass=admin666")
	w := request("POST", "/ZlsManage/UserApi/GetToken.go", &stBody{Body: body, ContentType: "application/x-www-form-urlencoded"})

	res := w.Body.String()

	tt.EqualExit(200, w.Code)
	tt.EqualExit(200, zjson.Get(res, "code").Int())
	tt.EqualTrue(zjson.Get(res, "data.token").Exists())

	t.Log("登录成功", res)
}
