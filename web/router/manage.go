package router

import (
	"app/global"
	"app/web/controller/manage"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/znet/cors"
	"github.com/sohaha/zlsgo/zutil"
)

// RegHome 注册 后台 路由
func (*StController) RegManage(r *znet.Engine) {
	if global.DB == nil {
		global.Log.Error("有没使用数据库，无法使用管理后台功能")
		return
	}

	r.Group("/ZlsManage/", func(r *znet.Engine) {
		corsHandler := cors.New(&cors.Config{
			Headers: []string{"Origin", "No-Cache", "X-Requested-With", "If-Modified-Since", "Pragma", "Last-Modified", "Cache-Control", "Expires", "Content-Type", "Access-Control-Allow-Origin", "token"},
		})
		r.Use(corsHandler)
		r.OPTIONS("*", func(c *znet.Context) {})

		r.Use(manage.Authority())

		r.BindStructDelimiter = ""
		r.BindStructSuffix = ".go"

		zutil.CheckErr(r.BindStruct("/UserApi", &manage.Basic{}))

		zutil.CheckErr(r.BindStruct("/UserManageApi", &manage.UserManage{}))

		r.Any("*", func(c *znet.Context) {
			c.ApiJSON(404, "404", nil)
		})
	})
}
