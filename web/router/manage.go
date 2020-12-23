package router

import (
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zstatic"

	"app/global"
	"app/web/controller/manage"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/znet/cors"
	"github.com/sohaha/zlsgo/znet/gzip"
	"github.com/sohaha/zlsgo/zutil"
)

// RegHome 注册 后台 路由
func (*StController) RegManage(r *znet.Engine) {
	if global.DB == nil {
		global.Log.Error("有没使用数据库，无法使用管理后台功能")
		return
	}

	g := gzip.Default()

	prefix := "manage"
	fileserver := zstatic.NewFileserver(global.ManageConf().Path)
	r.GET(prefix, func(c *znet.Context) {
		c.Redirect(prefix + "/")
	})
	r.GET(prefix+"/{file:.*}", fileserver, g)

	r.Static("/static/", zfile.RealPath("./resource/static"))

	r.Group("/ZlsManage/", func(r *znet.Engine) {
		corsHandler := cors.New(&cors.Config{
			Headers: []string{"Origin", "No-Cache", "X-Requested-With", "If-Modified-Since", "Pragma", "Last-Modified", "Cache-Control", "Expires", "Content-Type", "Access-Control-Allow-Origin", "token"},
		})
		r.Use(corsHandler, g)
		r.OPTIONS("*", func(c *znet.Context) {})

		r.Use(manage.Authority())

		r.BindStructDelimiter = ""
		r.BindStructSuffix = ".go"

		zutil.CheckErr(r.BindStruct("/UserApi", &manage.Basic{}))

		zutil.CheckErr(r.BindStruct("/UserManageApi", &manage.UserManage{}))

		zutil.CheckErr(r.BindStruct("/SystemApi", &manage.System{}))

		zutil.CheckErr(r.BindStruct("/MenuApi", &manage.Menu{}))

		r.Any("*", func(c *znet.Context) {
			c.ApiJSON(404, "404", nil)
		})
	})
}
