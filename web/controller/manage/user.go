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

// 后台-用户接口
type Basic struct{}

// PostGetToken 用户登录
func (*Basic) PostGetToken(c *znet.Context) {
	var user, token = model.AuthUser{}, ""
	var tokenID uint

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
	token, tokenID, err = user.Login(ip, ua)
	if err != nil {
		c.ApiJSON(212, err.Error(), nil)
		return
	}

	web.ApiJSON(c, 200, "登录成功", user, map[string]interface{}{
		"token": token,
		"tid":   tokenID,
	})
}

// GetUseriInfo 用户详情
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

	menu := (&model.AuthGroupMenu{GroupID: uint8(user.GroupID)}).GroupMenu(user)

	web.ApiJSON(c, 200, "用户详情", user, map[string]interface{}{
		"last":    t,
		"systems": systems,
		"groups":  groups,
		"marks":   (model.AuthUserGroup{ID: user.GroupID}).GetMarks(),
		"router":  (&model.AuthGroupMenu{}).MenuReg(menu),
		"menu":    menu,
	})
}

// PutUpdate 更新用户资料
func (*Basic) PutUpdate(c *znet.Context) {
	user := &model.AuthUser{}
	if u, ok := c.Value("user"); ok {
		user = u.(*model.AuthUser)
	}

	var postData manageBusiness.PutUpdateSt
	if err := c.Bind(&postData); err != nil {
		web.ApiJSON(c, 201, err.Error(), nil)
		return
	}

	uid := user.ID
	currentUserId := uid
	if postData.Id > 0 {
		currentUserId = postData.Id
	}

	// isAdmin := manageBusiness.IsAdmin(uid)
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

	// _, err := user.Update(c, postData, currentUserId, isAdmin, isMe)
	_, err := user.Update(c, postData, currentUserId, VerifPermissionMark(c, "password"))
	if err != nil {
		web.ApiJSON(c, 201, err.Error(), nil)
		return
	}

	web.ApiJSON(c, 200, "处理成功", map[string]int{"result": 1})
}

// PutEditPassword 修改用户密码
func (*Basic) PutEditPassword(c *znet.Context) {
	user := &model.AuthUser{}
	if u, ok := c.Value("user"); ok {
		user = u.(*model.AuthUser)
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

	userid := user.ID
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

// PostUploadAvatar 上传用户头像
func (*Basic) PostUploadAvatar(c *znet.Context) {
	// *multipart.FileHeader
	file, _ := c.FormFile("file")
	if file == nil {
		web.ApiJSON(c, 211, "请选择图片", nil)
		return
	}

	rePath, host, err := manageBusiness.UploadAvatar(file, c.Host())
	if err != nil {
		web.ApiJSON(c, 211, err.Error(), nil)
		return
	}

	web.ApiJSON(c, 200, "上传成功", map[string]interface{}{
		"path": rePath,
		"host": host,
	})
	return
}

// PostClearToken 清除用户Token
func (*Basic) PostClearToken(c *znet.Context) {
	t, ok := c.Value("token")
	if ok {
		t.(*model.AuthUserToken).UpdateStatus()
	}
	web.ApiJSON(c, 200, "退出完成", true)
}
