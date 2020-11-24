package main

import (
	"app/global"
	"app/model"

	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zvalid"
)

type passwdCli struct {
	User   *string
	Passwd *string
}

func (c *passwdCli) Flags(*zcli.Subcommand) {
	c.User = zcli.SetVar("u", "用户名").String()
	c.Passwd = zcli.SetVar("p", "用户新密码").String()
}

func (c *passwdCli) Run([]string) {
	username, err := zvalid.Text(*c.User, "用户名").Required().String()
	if err != nil {
		global.Log.Error(err)
		return
	}
	encryptPassword, err := zvalid.Text(*c.Passwd, "用户新密码").Required().EncryptPassword().String()
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.InitConf()
	u := &model.AuthUser{}

	global.DB.Model(u).Where(&model.AuthUser{Username: username}).Limit(1).Find(&u)
	if u.ID == 0 {
		global.Log.Error("用户不存在")
		return
	}

	tx := global.DB.Model(u).Select("password").Where("username = ?", username).Updates(&model.AuthUser{Password: encryptPassword})
	if tx.RowsAffected == 0 {
		global.Log.Error("密码更新失败")
		return
	}
	global.Log.Success("密码更新成功")
}
