package router

import (
	"app/compose"
	"app/web/middleware"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zpprof"
	"github.com/sohaha/zlsgo/zutil"
)

type StController struct {
}

// Engine 路由服务
var Engine *znet.Engine

func init() {
	Engine = znet.New()
}

// Run 启动服务
func Run() {
	znet.Run()
}

// Init 初始化路由
func Init() {
	// 复用日志对象
	Engine.Log = compose.Log

	// 设置开发模式
	if compose.BaseConf().Debug && compose.WebConf().Debug {
		Engine.SetMode(znet.DebugMode)
		// conf.Log.Discard()
	}

	// 性能分析
	if compose.WebConf().Pprof {
		zpprof.Register(Engine, compose.WebConf().PprofToken)
	}

	// 绑定端口
	webPort := compose.EnvPort
	if webPort == "" {
		webPort = compose.WebConf().Port
	}
	Engine.SetAddr(webPort)

	// 注册全局中间件
	middleware.RegisterMiddleware(Engine)

	// 注册路由
	registerController(Engine)

	// 未知路由处理
	Engine.NotFoundHandler(func(c *znet.Context) {
		if c.IsAjax() {
			c.ApiJSON(404, "NotFound", nil)
			return
		}
		c.String(404, "NotFound")
	})

	// 设置 HTTPS
	if compose.WebConf().Tls && compose.WebConf().TlsPort != "" {
		Engine.AddAddr(":"+compose.WebConf().TlsPort, znet.TlsCfg{
			// http 重定向 https
			HTTPAddr: webPort,
			Key:      compose.WebConf().Key,
			Cert:     compose.WebConf().Cert, // or domain.crt
		})
	}
}

func registerController(r *znet.Engine) {
	_ = zutil.RunAllMethod(&StController{}, r)
}
