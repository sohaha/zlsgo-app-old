package router

import (
	"app/web/controller"
	"app/web/controller/manage"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"
)

// RegHome 注册 Home 路由
func (*StController) RegHome(r *znet.Engine) {
	homeController := &controller.Home{}

	err := r.BindStruct("/", homeController)
	zutil.CheckErr(err)

	// r.GET("/", homeController.Home)
}

// RegHome 注册 Home 路由
func (*StController) RegManage(r *znet.Engine) {
	r.Group("/ZlsManage/", func(r *znet.Engine) {
		r.BindStructDelimiter = ""
		r.BindStructSuffix = ".go"

		var err = r.BindStruct(
			"/UserApi",
			&manage.User{},
			manage.Authority(),
		)
		zutil.CheckErr(err)
	})
}
