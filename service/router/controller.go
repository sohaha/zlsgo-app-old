package router

import (
	"app/service/api"
	"github.com/sohaha/zlsgo/znet"
)

// RegHome 注册路由
func (*ControllerSt) RegHome(r *znet.Engine) {
	homeController := &api.Home{}

	r.GET("/", homeController.Home)
}
