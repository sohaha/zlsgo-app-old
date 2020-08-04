package router

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"

	"app/web"
)

// RegHome 注册 Home 路由
func (*StController) RegHome(r *znet.Engine) {
	homeController := &web.Home{}

	err := r.BindStruct("/", homeController)
	zutil.CheckErr(err)
	// r.GET("/", homeController.Home)
}
