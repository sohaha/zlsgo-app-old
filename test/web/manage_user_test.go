package web_test

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"app/model"
	. "github.com/sohaha/zlsgo"
	"github.com/sohaha/zlsgo/zjson"
)

func testUser(t *testing.T) {
	tt := NewTest(t)
	res, json := &httptest.ResponseRecorder{}, zjson.Res{}

	// 用户名空
	res, json = userLogin(t, "", "111")
	tt.EqualExit(200, res.Code)
	tt.EqualTrue(json.Get("code").Int() != 200)

	// 密码错误
	res, json = userLogin(t, "manage", "111")
	tt.EqualTrue(json.Get("code").Int() != 200)

	// 登录成功
	res, json = userLogin(t, "manage", model.DefManagePassword)
	tt.EqualExit(200, json.Get("code").Int())
	tt.EqualTrue(json.Get("data.token").Exists())
	manageToken = json.Get("data.token").String()

	res, json = userInfo(t)
	tt.EqualExit(200, json.Get("code").Int())
}

func userLogin(t *testing.T, username, password string) (*httptest.ResponseRecorder, zjson.Res) {
	body := strings.NewReader(fmt.Sprintf("user=%s&pass=%s", username, password))
	w := request("POST", "/ZlsManage/UserApi/GetToken.go", &stBody{
		Body: body,
		Header: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	})
	res := w.Body.String()
	t.Log(zjson.Get(res, "msg").String())
	return w, zjson.Parse(res)
}

func userInfo(t *testing.T) (*httptest.ResponseRecorder, zjson.Res) {
	w := request("GET", "/ZlsManage/UserApi/UseriInfo.go", &stBody{
		Header: map[string]string{
			"token": manageToken,
		},
	})
	res := w.Body.String()
	t.Log(zjson.Get(res, "msg").String())
	return w, zjson.Parse(res)
}
