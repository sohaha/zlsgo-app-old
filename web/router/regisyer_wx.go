package router

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zutil"

	"app/web/controller/wx"
)

// RegWx 注册微信路由
func (*StController) RegWx(r *znet.Engine) {
	r = r.Group("/wx")

	err := r.BindStruct("/mp", &wx.Mp{})
	zutil.CheckErr(err)

	err = r.BindStruct("/qy", &wx.Qy{})
	zutil.CheckErr(err)

	err = r.BindStruct("/open", &wx.Open{})
	zutil.CheckErr(err)
}
