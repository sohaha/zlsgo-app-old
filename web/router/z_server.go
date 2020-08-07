package router

import (
	"app/module"
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
	Engine.Log = module.Log

	// 设置开发模式
	if module.BaseConf().Debug && module.WebConf().Debug {
		Engine.SetMode(znet.DebugMode)
		// conf.Log.Discard()
	}

	// 性能分析
	if module.WebConf().Pprof {
		zpprof.Register(Engine, module.WebConf().PprofToken)
	}

	// 绑定端口
	webPort := module.EnvPort
	if webPort == "" {
		webPort = module.WebConf().Port
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
	if module.WebConf().Tls && module.WebConf().TlsPort != "" {
		Engine.AddAddr(":"+module.WebConf().TlsPort, znet.TlsCfg{
			// http 重定向 https
			HTTPAddr: webPort,
			Key:      module.WebConf().Key,
			Cert:     module.WebConf().Cert, // or domain.crt
		})
	}
}

func registerController(r *znet.Engine) {
	_ = zutil.RunAllMethod(&StController{}, r)
}
