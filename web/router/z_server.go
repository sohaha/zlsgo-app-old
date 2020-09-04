package router

import (
	"app/global"
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
	znet.ShutdownDone = func() {
		// 设置程序关闭前回收操作
		global.Recover()
	}
	znet.Run()
}

// Init 初始化路由
func Init() {
	// 复用日志对象
	Engine.Log = global.Log

	// 设置开发模式
	if global.BaseConf().Debug && global.WebConf().Debug {
		Engine.SetMode(znet.DebugMode)
		// conf.Log.Discard()
	}

	// 性能分析
	if global.WebConf().Pprof {
		zpprof.Register(Engine, global.WebConf().PprofToken)
	}

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

	// 绑定端口
	webPort := global.EnvPort
	if webPort == "" {
		webPort = global.WebConf().Port
	}

	// 设置 HTTPS
	if global.WebConf().Tls && global.WebConf().TlsPort != "" {
		Engine.SetAddr(global.WebConf().TlsPort, znet.TlsCfg{
			// http 重定向 https
			HTTPAddr: webPort,
			Key:      global.WebConf().Key,
			Cert:     global.WebConf().Cert, // or domain.crt
		})
	} else {
		Engine.SetAddr(webPort)
	}
}

func registerController(r *znet.Engine) {
	_ = zutil.RunAllMethod(&StController{}, r)
}
