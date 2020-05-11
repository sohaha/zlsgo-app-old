package router

import (
	"app/conf"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zpprof"
	"github.com/sohaha/zlsgo/zutil"
)

type ControllerSt struct {
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

func Closed(fn func())  {
	znet.ShutdownDone = fn
}

// Init 初始化路由
func Init() {
	// 复用日志对象
	Engine.Log = conf.Log

	// 设置开发模式
	if conf.Base().Debug && conf.Web().Debug {
		Engine.SetMode(znet.DebugMode)
		// conf.Log.Discard()
	}

	// 性能分析
	if conf.Web().Pprof {
		zpprof.Register(Engine, conf.Web().PprofToken)
	}

	// 绑定端口
	webPort := conf.EnvPort
	if webPort == "" {
		webPort = conf.Web().Port
	}
	Engine.SetAddr(webPort)

	// 注册全局中间件
	registerMiddleware(Engine)

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
	if conf.Web().Tls && conf.Web().TlsPort != "" {
		Engine.AddAddr(":"+conf.Web().TlsPort, znet.TlsCfg{
			// http 重定向 https
			HTTPAddr: webPort,
			Key:      conf.Web().Key,
			Cert:     conf.Web().Cert, // or domain.crt
		})
	}
}

func registerController(r *znet.Engine) {
	_ = zutil.RunAllMethod(&ControllerSt{}, r)
}
