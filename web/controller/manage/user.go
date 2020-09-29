package manage

import (
	"app/web"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"

	"app/model"
)

type Basic struct{}

func (*Basic) PostGetToken(c *znet.Context) {
	var user, token = model.AuthUser{}, ""

	v := c.ValidRule()
	err := zvalid.Batch(
		zvalid.BatchVar(&user.Username, v.Verifi(c.DefaultFormOrQuery("user", ""), "用户名").Required()),
		zvalid.BatchVar(&user.Password, v.Verifi(c.DefaultFormOrQuery("pass", ""), "用户密码").Required()),
	)
	ip := c.GetClientIP()
	ua := c.GetHeader("User-Agent")
	token, err = user.Login(ip, ua)
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	web.ApiJSON(c, 200, "Done", user, map[string]interface{}{
		"token": token,
	})
}

func (*Basic) GetUseriInfo(c *znet.Context) {
	user := &model.AuthUser{}
	if u, ok := c.Value("user"); ok {
		user = u.(*model.AuthUser)
	}
	t := &model.AuthUserToken{
		Userid: user.ID,
	}
	t.Last()
	web.ApiJSON(c, 200, "Done", user, map[string]interface{}{
		"last":    t,
		"systems": map[string]interface{}{},
	})
}

func (*Basic) PutUpdate(c *znet.Context) {

}

func (*Basic) PutEditPassword(c *znet.Context) {

}

func (*Basic) PostUploadAvatar(c *znet.Context) {

}

func (*Basic) PostClearToken(c *znet.Context) {
	t, ok := c.Value("token")
	if ok {
		t.(*model.AuthUserToken).UpdateStatus()
	}
	web.ApiJSON(c, 200, "退出完成", nil)
}
