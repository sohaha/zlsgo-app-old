package manage

import (
	"app/logic"
	"app/web"
	"app/web/business/manageBusiness"
	"errors"
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
	u, ok := c.Value("user")
	if !ok {
		web.ApiJSON(c, 212, "请登录", nil)
	}

	var postData manageBusiness.PutUpdateSt
	if err := c.Bind(&postData); err != nil {
		web.ApiJSON(c, 201, err.Error(), nil)
		return
	}

	uid := u.(*model.AuthUser).ID
	currentUserId := uid
	if postData.Id > 0 {
		currentUserId = postData.Id
	}

	isAdmin := manageBusiness.IsAdmin(uid)
	userIsAdmin := manageBusiness.IsAdmin(currentUserId)
	isMe := currentUserId == uid

	if isMe && 1 != postData.Status {
		web.ApiJSON(c, 201, "不能禁止自己", nil)
		return
	}

	if (currentUserId != uid) && (userIsAdmin == 1) && (postData.Status != 0) {
		web.ApiJSON(c, 201, "不能更新该账户状态", nil)
		return
	}

	_, err := u.(*model.AuthUser).Update(c, postData, currentUserId, isAdmin, isMe)
	if err != nil {
		web.ApiJSON(c, 201, err.Error(), nil)
		return
	}

	web.ApiJSON(c, 200, "处理成功", map[string]int{"result": 1})
}

func (*Basic) PutEditPassword(c *znet.Context) {
	u, ok := c.Value("user")
	if !ok {
		web.ApiJSON(c, 212, "请登录", nil)
	}

	var postData manageBusiness.PutEditPasswordSt
	valid := c.ValidRule()
	err := c.BindValid(&postData, map[string]zvalid.Engine{
		"oldPass": valid.Required("请输入旧密码"),
		"pass": valid.Required("请输入新密码").MinLength(3, "密码最少3字符").MaxLength(50, "密码最多50字符").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			if rawValue != c.DefaultPostForm("pass2", "") {
				newErr = errors.New("两次密码不一致")
			}
			newValue = rawValue
			return
		}).EncryptPassword(),
	})

	if err != nil {
		web.ApiJSON(c, 201, err.Error(), nil)
		return
	}
	
	userid := u.(*model.AuthUser).ID
	upUid := postData.UserID
	if upUid == 0 {
		upUid = userid
	}
	if userid == upUid {
		if err := (&model.AuthUser{ID: upUid}).EditPassword(c, postData); err != nil {
			web.ApiJSON(c, 201, err.Error(), nil)
			return
		}
		_ = (&model.AuthUserToken{Userid: upUid}).ClearAllToken()

		web.ApiJSON(c, 200, "修改密码成功", nil)
		return
	} else {
		web.ApiJSON(c, 201, "不能修改其他人密码", nil)
		return
	}
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
