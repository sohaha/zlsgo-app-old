package manage

import (
	"os"

	"app/dal/model"
	"app/dal/query"
	"app/global"
	"app/logic"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztime"
	"gorm.io/gorm"
)

type Base struct{}

// Get 首页
func (*Base) Get(c *znet.Context) {
	c.ApiJSON(200, "服务正常", map[string]interface{}{
		"time": ztime.Now(),
		"pid":  os.Getpid(),
	})
}

func (*Base) PostLogin(c *znet.Context) {
	j, err := c.GetJSONs()
	if err != nil {
		return
	}

	username := j.Get("user").String()
	passwd := j.Get("pass").String()

	db := global.DB
	u := query.Use(db).ManageUser
	user, err := u.Where(u.Username.Eq(username)).Select(u.Username, u.Avatar, u.Password).Last()
	if err != nil {
		errmsg := err.Error()
		if err == gorm.ErrRecordNotFound {
			errmsg = "用户不存在/密码错误"
		}
		c.ApiJSON(211, errmsg, nil)
		return
	}

	if !logic.PasswordVerify(passwd, user.Password) {
		c.ApiJSON(211, "用户不存在/密码错误", nil)
		return
	}

	type PublicUser struct {
		*model.ManageUser
		Password  *struct{} `json:"password,omitempty"`
		ID        *struct{} `json:"id,omitempty"`
		Key       *struct{} `json:"key,omitempty"`
		Effective *struct{} `json:"effective,omitempty"`
	}

	c.ApiJSON(200, "服务正常", PublicUser{
		ManageUser: user,
	})
}
