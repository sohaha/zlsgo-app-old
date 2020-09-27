package router

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"

	"app/web/controller/manage"
)

// RegHome 注册 后台 路由
func (*StController) RegManage(r *znet.Engine) {

	r.Group("/ZlsManage/", func(r *znet.Engine) {
		r.Use(manage.Authority())

		r.BindStructDelimiter = ""
		r.BindStructSuffix = ".go"
		zutil.CheckErr(r.BindStruct("/UserApi", &manage.User{}))

		zutil.CheckErr(r.BindStruct("/UserManageApi", &manage.UserManage{}))

	})
}
