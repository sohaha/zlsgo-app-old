package manage

import (
	"app/logic"
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
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	ip := c.GetClientIP()
	ua := c.GetHeader("User-Agent")
	token, err = user.Login(ip, ua)
	if err != nil {
		c.ApiJSON(212, err.Error(), nil)
		return
	}

	web.ApiJSON(c, 200, "登录成功", user, map[string]interface{}{
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
	// 上次登录信息
	t.Last()

	systems := map[string]interface{}{}
	// 有系统管理权限需要打印系统信息
	if VerifPermissionMark(c, "systems", true) {
		systems = logic.GetServerInfo()
	}

	groups := []model.AuthUserGroup{}
	(model.AuthUserGroup{}).All(&groups)

	web.ApiJSON(c, 200, "用户详情", user, map[string]interface{}{
		"last":    t,
		"systems": systems,
		"groups":  groups,
		"marks":   (model.AuthUserGroup{ID: user.GroupID}).GetMarks(),
		// menus
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
