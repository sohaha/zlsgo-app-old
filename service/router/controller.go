package router

import (
	"app/service/api"
	"app/service/api/wx"
	"github.com/sohaha/zlsgo/znet"
)

// RegHome 注册路由
func (*ControllerSt) RegHome(r *znet.Engine) {
	homeController := &api.Home{}

	r.GET("/", homeController.Home)
	r.GET("/token", homeController.Token)
}

func (*ControllerSt) RegWx(r *znet.Engine) {
	wxMp := &wx.Mp{}
	wxQy := &wx.Qy{}
	wxOpen := &wx.Open{}

	g := r.Group("/wx")

	mp := g.Group("/mp")
	{
		mp.GET("/token", wxMp.AccessToken)
		mp.GET("/jsapiTicket", wxMp.JsapiTicket)
	}

	qy := g.Group("/qy")
	{
		qy.GET("/token", wxQy.AccessToken)
		qy.GET("/jsapiTicket", wxQy.JsapiTicket)
	}

	open := g.Group("/open")
	{
		open.POST("/notification", wxOpen.Notification)
		open.GET("/ticket", wxOpen.Ticket)
		open.GET("/apiQueryAuth", wxOpen.ApiQueryAuth)
		open.GET("/token", wxOpen.AccessToken)
		open.GET("/jsapiTicket", wxOpen.JsapiTicket)
	}

}
