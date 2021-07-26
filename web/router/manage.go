package router

import (
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zstatic"
	"github.com/zlsgo/resource"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/znet/cors"
	"github.com/sohaha/zlsgo/znet/gzip"
	"github.com/sohaha/zlsgo/zutil"

	"app/global"
	"app/web/controller/manage"
)

// RegHome 注册 后台 路由
func (*StController) RegManage(r *znet.Engine) {
	prefix := "manage"

	if global.DB == nil {
		global.Log.Error("有没使用数据库，无法使用管理后台功能")
		return
	}

	g := gzip.Default()

	// 静态资源目录，常用于放上传的文件
	r.Static("/static/", zfile.RealPathMkdir("./resource/static"))

	// 注意： 这里的路径不能直接使用变量 global.ManageConf().Path ，但需要保持一致
	fileServ, group := zstatic.NewFileserverAndGroup("resource/manage")
	if _, e := group.MustBytes("index.html"); e != nil {
		// 初始化后台资源
		initManageResource("resource/manage")
		return
	}
	r.GET(prefix+"/{file:.*}", fileServ, g)

	// 后台
	// r.Static("/manage/", zfile.RealPathMkdir("./resource/manage"))

	r.Group("/Manage/", func(r *znet.Engine) {
		// 开启跨域
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

func initManageResource(path string) {
	c := global.ManageConf()
	var err error
	defer func() {
		if err != nil {
			global.Log.Error("Manage init failed")
		} else {
			global.Log.Success("Manage init complete")
		}
	}()
	r := resource.New(c.Remote)
	r.SetMd5(c.Md5)
	r.SetDeCompressPath(path)
	r.SetFilterRule([]string{"(.*)/\\.git/", "(.*)/\\.vscode/", "(.*)/\\.idea/"})
	err = r.SilentRun(func(current, total int64) {
	})
}
