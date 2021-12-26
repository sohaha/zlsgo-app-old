package router

import (
	"app/web"
	"app/web/controller/manage"
	"app/web/controller/manage/permission"
	"github.com/sohaha/zlsgo/znet"
)

// RegManage Manage路由
func (*StController) RegManage(r *znet.Engine) {
	g := r.Group("manage/")
	{
		g.BindStructDelimiter = "_"
		g.BindStructSuffix = ".go"

		g.Use(func(c *znet.Context) {
			c.Next()
		})

		g.Any("*", func(c *znet.Context) {
			web.ApiJSON(c, 404, "此路不通", nil)
		})

		_ = g.BindStruct("base", &manage.Base{})
		_ = g.BindStruct("user", &manage.User{})
		g = g.Group("permission/")

		_ = g.BindStruct("role", &permission.Role{})
		_ = g.BindStruct("rule", &permission.Rule{})
	}
}
