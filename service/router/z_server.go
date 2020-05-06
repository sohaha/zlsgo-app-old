package router

import (
	"strings"

	"app/conf"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zpprof"
	"github.com/sohaha/zlsgo/zutil"
)

type ControllerSt struct {
}

func Run() {
	// 初始化服务
	r := znet.New()

	// 复用日志对象
	r.Log = conf.Log

	// 设置开发模式
	if conf.Base().Debug && conf.Web().Debug {
		r.SetMode(znet.DebugMode)
		// conf.Log.Discard()
	}

	// 性能分析
	if conf.Web().Pprof {
		zpprof.Register(r, conf.Web().PprofToken)
	}

	// 绑定端口
	webPort := conf.EnvPort
	if webPort == "" {
		webPort = conf.Web().Port
	}
	r.SetAddr(webPort)

	// 未知路由处理
	r.NotFoundHandler(func(c *znet.Context) {
		if c.IsAjax() {
			c.ApiJSON(404, "NotFound", nil)
			return
		}
		c.String(404, "NotFound")
	})

	// 异常处理
	r.PanicHandler(func(c *znet.Context, err error) {
		if c.Engine.IsDebug() {
			errData := zlog.TrackCurrent(10, 4)
			if c.IsAjax() {
				c.ApiJSON(500, "Panic", errData)
				return
			}
			c.HTML(500, strings.Join(errData, "<br><br>"))
			return
		}
		if c.IsAjax() {
			c.ApiJSON(500, "Panic", nil)
			return
		}
		c.String(500, "Panic")
	})

	// 注册全局中间件
	registerMiddleware(r)

	// 注册路由
	registerController(r)

	// 设置 HTTPS
	if conf.Web().Tls && conf.Web().TlsPort != "" {
		r.AddAddr(":"+conf.Web().TlsPort, znet.TlsCfg{
			// http 重定向 https
			HTTPAddr: webPort,
			Key:      conf.Web().Key,
			Cert:     conf.Web().Cert, // or domain.crt
		})
	}

	// 启动服务
	znet.Run()
}

func registerController(r *znet.Engine) {
	_ = zutil.RunAllMethod(&ControllerSt{}, r)
}
