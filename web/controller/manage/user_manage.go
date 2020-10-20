package manage

import (
	"app/web"
	"app/web/business/manageBusiness"
	"errors"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"
	"strconv"

	"app/model"
)

type UserManage struct {
}

// GetUserLists 获取用户列表
func (*UserManage) GetUserLists(c *znet.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultFormOrQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultFormOrQuery("page", "1"))
	p := model.Page{
		Curpage:  uint(page),
		Pagesize: uint(pagesize),
	}
	users := (&model.AuthUser{}).Lists(&p)

	c.ApiJSON(200, "用户列表", map[string]interface{}{
		"items": users,
		"page":  p,
	})
}

// PostUser 创建用户
func (*UserManage) PostUser(c *znet.Context) {
	var user struct {
		model.AuthUser
		Password string `json:"password"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&user, map[string]zvalid.Engine{
		"username": valid.Required("用户名不能为空").MaxUTF8Length(20, "用户名最多20字符"),
		"password": valid.Required().Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if rawValue != c.DefaultPostForm("password2", "") {
				newErr = errors.New("两次密码不一致")
			}
			newValue = rawValue
			return
		}).EncryptPassword(),
		"status": valid.EnumInt([]int{1, 2}, "用户状态值错误"),
		"email":  valid.Required("Email地址不能为空").IsMail("Email地址错误"),
		"remark": valid.MaxLength(200, "用户简介最多200字符"),
		"avatar": valid.MaxLength(250, "头像地址不能超过250字符"),
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
// php上传参格式 x-www-form-urlencoded
// go上传参格式 form-data
func (*UserManage) DeleteUser(c *znet.Context) {
	u, ok := c.Value("user")
	if !ok {
		web.ApiJSON(c, 212, "请登录", nil)
	}

	id, err := c.Valid(zvalid.New(), "id", "id").Required("id不能为空").Int()
	if err != nil || id == 0 {
		c.ApiJSON(200, "删除用户", 0)
		return
	}

	switch true {
	case uint(id) == u.(*model.AuthUser).ID:
		err = errors.New("不可以删除自己")
		break
	case manageBusiness.IsAdmin(uint(id)) == 1:
		err = errors.New("请移除该用户的超级管理员身份")
		break
	}

	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	if err := (&model.AuthUser{ID: uint(id)}).Delete(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "删除用户", 1)
	return
}

// GetGroups 获取角色列表
func (*UserManage) GetGroups(c *znet.Context) {
	var res []model.AuthUserGroup
	(&model.AuthUserGroup{}).All(&res)

	c.ApiJSON(200, "角色列表", res)
	return
}

// GetGroupInfo 获取角色详情
func (*UserManage) GetGroupInfo(c *znet.Context) {
	id, err := c.Valid(zvalid.New(), "id", "id").Required("id不能为空").Int()
	if err != nil || id == 0 {
		c.ApiJSON(211, "参数错误", nil)
		return
	}

	res := &model.AuthUserGroup{ID: uint(id)}
	res.GroupInfo()
	integration := (&model.AuthUserRulesRela{GroupID: res.ID}).Integration()

	web.ApiJSON(c, 200, "角色详情", res, map[string]interface{}{
		"rule_ids":     integration.RuleIds,
		"ban_rule_ids": integration.BanRuleIds,
		"user_count":   integration.UserCount,
	})
}

// 创建角色
func (*UserManage) PostGroups(c *znet.Context) {
	var postParam struct {
		Name   string `json:"name"`
		Remark string `json:"remark"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"name": valid.Required("角色名称不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			if err := (&model.AuthUserGroup{Name: rawValue}).Exist(); err != nil {
				newErr = err
				return
			}

			newValue = rawValue
			return
		}),
		"remark": valid.Customize(func(rawValue string, err error) (newValue string, newErr error) {
			newValue = rawValue
			return
		}),
	})

	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	res := &model.AuthUserGroup{Name: postParam.Name, Remark: postParam.Remark}
	if err := res.Save(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "创建新角色", res)
	return
}

// 更新角色
func (*UserManage) PutGroups(c *znet.Context) {
	var postParam struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Remark string `json:"remark"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"id": valid.Customize(func(rawValue string, err error) (newValue string, newErr error) {
			newValue = rawValue
			return
		}),
		"name": valid.Required("角色名称不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			if err := (&model.AuthUserGroup{Name: rawValue, ID: postParam.ID}).Exist(); err != nil {
				newErr = err
				return
			}

			newValue = rawValue
			return
		}),
		"remark": valid.Customize(func(rawValue string, err error) (newValue string, newErr error) {
			newValue = rawValue
			return
		}),
	})

	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	res := &model.AuthUserGroup{Name: postParam.Name, Remark: postParam.Remark, ID: postParam.ID}
	if err := res.Save(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "更新角色", res)
	return
}
