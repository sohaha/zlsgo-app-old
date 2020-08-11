package router

import (
	"app/web/controller"
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
