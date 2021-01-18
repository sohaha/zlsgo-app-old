package manage

import (
	"app/model"
	"encoding/json"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"
	"strconv"
)

// 后台-用户管理接口
type Menu struct {
}

// PostUserMenu 获取全部菜单
func (*Menu) PostUserMenu(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	groupid, err := strconv.Atoi(c.GetJSON("groupid").String())
	if err != nil {
		groupid = 0
	}

	c.ApiJSON(200, "请求成功", (&model.Menu{}).Lists(uint8(groupid)))
	return
}

// PostCreate 新增菜单
func (*Menu) PostCreate(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var postParam struct {
		Title      string `json:"title"`
		Index      string `json:"index"`
		Icon       string `json:"icon"`
		Breadcrumb int    `json:"breadcrumb"`
		Real       int    `json:"real"`
		Show       int    `json:"show"`
		Pid        int    `json:"pid"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"title": valid.Trim().Required("菜单名称不能为空"),
		"index": valid.Trim().Required("路由不能为空"),
	})
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	if postParam.Pid != 0 {
		hasInfo := &model.Menu{ID: uint(postParam.Pid)}
		if hasInfo.Exist(); hasInfo.Title == "" {
			c.ApiJSON(211, "pid不合法", nil)
			return
		}
	}

	res := &model.Menu{
		Title:      postParam.Title,
		Index:      postParam.Index,
		Icon:       postParam.Icon,
		Breadcrumb: uint8(postParam.Breadcrumb),
		Real:       uint8(postParam.Real),
		Show:       uint8(postParam.Show),
		Pid:        uint8(postParam.Pid),
	}
	if err := res.Create(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "处理成功", map[string]interface{}{"id": res.ID})
	return
}

// PostDelete 删除菜单
func (*Menu) PostDelete(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	id, err := strconv.Atoi(c.GetJSON("id").String())
	if err != nil || id == 0 {
		c.ApiJSON(211, "菜单id不允许为空", nil)
		return
	}

	hasInfo := &model.Menu{Pid: uint8(id)}
	if hasInfo.PidExist(); hasInfo.Title != "" {
		c.ApiJSON(211, "请先删除子集", nil)
		return
	}

	if err := (&model.Menu{ID: uint(id)}).Delete(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "处理成功", nil)
	return
}

// PostUpdate 更新菜单
func (*Menu) PostUpdate(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var postParam struct {
		ID         int    `json:"id"`
		Title      string `json:"title"`
		Index      string `json:"index"`
		Icon       string `json:"icon"`
		Breadcrumb int    `json:"breadcrumb"`
		Real       int    `json:"real"`
		Show       int    `json:"show"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"id":    valid.Required("菜单id不允许为空"),
		"title": valid.Trim().Required("菜单名称不能为空"),
		"index": valid.Trim().Required("路由不能为空"),
	})
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	res := &model.Menu{
		ID:         uint(postParam.ID),
		Title:      postParam.Title,
		Index:      postParam.Index,
		Icon:       postParam.Icon,
		Breadcrumb: uint8(postParam.Breadcrumb),
		Real:       uint8(postParam.Real),
		Show:       uint8(postParam.Show),
	}
	if err := res.Update(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "处理成功", nil)
	return
}

// PostSort 菜单拖拽排序(支持多次拖拽一起排)
func (*Menu) PostSort(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	menu := c.DefaultFormOrQuery("menu", "")
	var postSortMap model.PostSortSt
	json.Unmarshal([]byte(menu), &postSortMap)

	if err := (&model.Menu{}).MenuSort(postSortMap); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "处理成功", 1)
	return
}

// PostUpdateGroupMenu 角色菜单更新
func (*Menu) PostUpdateGroupMenu(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var postParam struct {
		GroupID int    `json:"groupid"`
		Menu    string `json:"menu"`
	}

	valid := c.ValidRule()
	err := c.BindValid(&postParam, map[string]zvalid.Engine{
		"groupid": valid.Required("角色id不能为空"),
		"menu":    valid.Trim().Required("菜单内容不能为空"),
	})
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	if err := (&model.AuthGroupMenu{GroupID: uint8(postParam.GroupID), Menu: postParam.Menu}).Update(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "处理成功", 1)
	return
}
