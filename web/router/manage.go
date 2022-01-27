package router

import (
	"app/web"
	"app/web/controller/manage"
	"app/web/controller/manage/permission"
	"app/web/middleware"
	"github.com/sohaha/zlsgo/znet"
)

// RegManage Manage路由
func (*StController) RegManage(r *znet.Engine) {
	g := r.Group("manage/")
	{
		g.BindStructDelimiter = "_"
		g.BindStructSuffix = ".go"

		pubRouters := []string{"/manage/base/login.go"}
		g.Use(middleware.Manage(pubRouters))
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
