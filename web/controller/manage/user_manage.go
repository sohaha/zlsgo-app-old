package manage

import (
	"errors"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"

	"app/model"
)

type UserManage struct {
}

// PostUser 创建用户
func (*UserManage) PostUser(c *znet.Context) {
	var user model.AuthUser

	valid := c.ValidRule()
	err := c.BindValid(&user, map[string]zvalid.Engine{
		"username": valid.Required(),
		"password": valid.Required().Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if rawValue != c.DefaultPostForm("password2", "") {
				newErr = errors.New("两次密码不一致")
			}
			newValue = rawValue
			return
		}).EncryptPassword(),
		"status": valid.Trim(),
	})

	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	err = user.Insert()
	if err != nil {
		c.ApiJSON(221, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "创建用户完成", user.ID)
}

// DeleteUser 删除用户
func (*UserManage) DeleteUser(c *znet.Context) {

}
