package manage

import (
	"app/logic"
	"app/web"
	"errors"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"
	"strconv"

	"app/model"
)

// 后台-用户管理接口
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
	user, authUserQuery := &model.AuthUser{}, &model.AuthUser{}
	if key := c.DefaultFormOrQuery("key", ""); key != "" {
		authUserQuery.Username = key
	}

	if u, ok := c.Value("user"); ok {
		user = u.(*model.AuthUser)
	}

	groupsHas1 := false
	for _, groupID := range user.GroupID {
		if groupID == 1 {
			groupsHas1 = true
		}
	}
	if !user.IsSuper && !groupsHas1 {
		authUserQuery.ID = user.ID
	}
	users := authUserQuery.Lists(&p)
	lists := (&model.AuthUser{}).ListsSub(users)

	c.ApiJSON(200, "用户列表", map[string]interface{}{
		"items": lists,
		"page":  p,
	})
}

// PostUser 创建用户
func (*UserManage) PostUser(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var user struct {
		model.AuthUser
		Password string `json:"password"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&user, map[string]zvalid.Engine{
		"username": valid.Required("用户名不能为空").MaxUTF8Length(20, "用户名最多20字符"),
		"password": valid.Required().Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if rawValue != c.GetJSON("password2").Str {
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

	// re get group_id => arr
	var groupIdKV []string
	c.GetJSON("group_id").ForEach(func(key, val zjson.Res) bool {
		groupIdKV = append(groupIdKV, val.String())
		return true
	})

	for _, groupID := range groupIdKV {
		if g, err := strconv.Atoi(groupID); err == nil {
			user.GroupID = append(user.GroupID, uint(g))
		}
	}

	err = (&user).Insert(user.Password)
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
	if !VerifPermissionMark(c, "systems") {
		return
	}

	user := &model.AuthUser{}
	if u, ok := c.Value("user"); ok {
		user = u.(*model.AuthUser)
	}

	id, err := c.Valid(zvalid.New(), "id", "id").Required("id不能为空").Int()
	if err != nil || id == 0 {
		c.ApiJSON(200, "删除用户", 0)
		return
	}

	switch true {
	case uint(id) == user.ID:
		err = errors.New("不可以删除自己")
		break
	case logic.IsAdmin(uint(id)) == 1:
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
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var res []model.AuthUserGroup
	(&model.AuthUserGroup{}).All(&res)

	c.ApiJSON(200, "角色列表", res)
	return
}

// GetGroupInfo 获取角色详情
func (*UserManage) GetGroupInfo(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

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

// PostGroups 创建角色
func (*UserManage) PostGroups(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

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

// PutGroups 更新角色
func (*UserManage) PutGroups(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

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

// DeleteGroups 删除角色
func (*UserManage) DeleteGroups(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	c.ApiJSON(211, "暂不支持", nil)
	return
}

// GetRules 获取权限规则列表
func (*UserManage) GetRules(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	key, _ := c.Valid(c.ValidRule(), "key", "key").String()
	res := (&model.AuthUserRules{Title: key, Mark: key}).Lists()

	resInterface := []map[string]interface{}{}
	for _, v := range res {
		resInterface = append(resInterface, map[string]interface{}{
			"id":     v.ID,
			"title":  v.Title,
			"mark":   v.Mark,
			"status": v.Status,
			"type":   v.Type,
			"remark": v.Remark,
		})
	}

	c.ApiJSON(200, "权限规则列表", resInterface)
	return
}

// PostRules 添加权限规则
func (*UserManage) PostRules(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var postParam struct {
		Title  string `json:"title"`
		Mark   string `json:"mark"`
		Type   uint8  `json:"type"`
		Remark string `json:"remark"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"title": valid.Required("名称不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			newValue = rawValue
			return
		}),
		"mark": valid.Required("标识不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}

			newValue = rawValue
			if err := (&model.AuthUserRules{Mark: newValue}).MarkExist(); err != nil {
				newErr = err
				return
			}

			return
		}),
		"type": valid.Required("类型不能为空").IsNumber("类型不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
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

	res := &model.AuthUserRules{Title: postParam.Title, Mark: postParam.Mark, Type: postParam.Type, Remark: postParam.Remark}
	if err := (res).Insert(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "权限规则列表", map[string]interface{}{"id": res.ID})
	return
}

// PutRules 编辑权限规则
func (*UserManage) PutRules(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var postParam struct {
		ID     uint   `json:"id"`
		Title  string `json:"title"`
		Mark   string `json:"mark"`
		Remark string `json:"remark"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"id": valid.Customize(func(rawValue string, err error) (newValue string, newErr error) {
			newValue = rawValue
			return
		}),
		"title": valid.Required("名称不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			newValue = rawValue
			return
		}),
		"mark": valid.Required("标识不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}

			newValue = rawValue
			if err := (&model.AuthUserRules{Mark: newValue, ID: postParam.ID}).MarkExist(); err != nil {
				newErr = err
				return
			}

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

	res := &model.AuthUserRules{ID: postParam.ID, Title: postParam.Title, Mark: postParam.Mark, Remark: postParam.Remark}
	if err := res.Update(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "权限规则列表", map[string]interface{}{
		"mark":   res.Mark,
		"remark": res.Remark,
		"res":    1.,
		"title":  res.Title,
	})
	return
}

// DeleteRules 删除权限规则
func (*UserManage) DeleteRules(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	id, err := c.Valid(c.ValidRule().Required(), "id", "id").Int()
	if err != nil {
		c.ApiJSON(200, "删除权限规则", 0)
		return
	}

	if err := (&model.AuthUserRules{ID: uint(id)}).Delete(); err != nil {
		c.ApiJSON(200, "删除权限规则", 0)
		return
	}

	c.ApiJSON(200, "删除权限规则", 1)
	return
}

// PutUpdateUserRuleStatus 更新用户规则权限
func (*UserManage) PutUpdateUserRuleStatus(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var postParam struct {
		ID     uint   `json:"id"`
		Gid    uint   `json:"gid"`
		Status uint8  `json:"status"`
		Sort   uint16 `json:"sort"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"id": valid.Required("角色权限不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			newValue = rawValue
			return
		}),
		"status": valid.Required("状态不能为空").EnumInt([]int{model.RelaStatusNormal, model.RelaStatusBan, model.RelaStatusIgnore}, "非正常状态码").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			newValue = rawValue
			return
		}),
		"sort": valid.MinInt(0, "排序范围为0-255").MaxInt(255, "排序范围为0-255").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			newValue = rawValue
			return
		}),
		"gid": valid.Required("用户组不能为空").Customize(func(rawValue string, err error) (newValue string, newErr error) {
			if err != nil {
				newErr = err
				return
			}
			newValue = rawValue
			return
		}),
	})

	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	res := &model.AuthUserRulesRela{RuleID: postParam.ID, GroupID: postParam.Gid, Status: postParam.Status, Sort: postParam.Sort}
	has, err := res.UpdateUserRuleStatus()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	re := map[string]interface{}{}
	if has.ID == 0 {
		re = map[string]interface{}{
			"group_id":    res.GroupID,
			"rule_id":     res.RuleID,
			"status":      res.Status,
			"sort":        res.Sort,
			"update_time": res.UpdatedAt,
			"create_time": res.CreatedAt,
		}
	} else {
		re = map[string]interface{}{
			"status":      res.Status,
			"sort":        res.Sort,
			"update_time": res.UpdatedAt,
		}
	}

	c.ApiJSON(200, "权限规则列表", []interface{}{
		re,
	})
	return
}
