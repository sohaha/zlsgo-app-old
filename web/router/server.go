package router

import (
	"app/conf"
	"app/global"

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

	baseConf := conf.Base()
	webConf := conf.Web()

	// 设置开发模式
	if baseConf.Debug && webConf.Debug {
		Engine.SetMode(znet.DebugMode)
	}

	// 性能分析
	if webConf.Pprof {
		zpprof.Register(Engine, webConf.PprofToken)
	}

	// 注册全局中间件
	// middleware.RegisterMiddleware(Engine)

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

	// 处理 panic
	Engine.PanicHandler(func(c *znet.Context, err error) {
		errMsg := ""
		if Engine.IsDebug() {
			errMsg = err.Error()
		}
		Engine.Log.Stack(errMsg)
		if c.IsAjax() {
			c.ApiJSON(500, "panic", errMsg)
			return
		}
		c.String(500, errMsg)
	})

	// 绑定端口
	webPort := conf.EnvPort
	if webPort == "" {
		webPort = webConf.Port
	}

	// 设置 HTTPS
	if webConf.Tls && webConf.TlsPort != "" {
		Engine.SetAddr(webConf.TlsPort, znet.TlsCfg{
			// http 重定向 https
			HTTPAddr: webPort,
			Key:      webConf.Key,
			Cert:     webConf.Cert, // or domain.crt
		})
	} else {
		Engine.SetAddr(webPort)
	}
}

func registerController(r *znet.Engine) {
	_ = zutil.RunAllMethod(&StController{}, r)
}
