package router

import (
	"strings"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/znet"
	_ "github.com/sohaha/zlsgo/znet/limiter"
)

func registerMiddleware(r *znet.Engine) {
	// 异常处理
	r.Use(znet.Recovery(r, func(c *znet.Context, err error) {
		if c.Engine.IsDebug() {
			errData := zlog.TrackCurrent(10, 4)
			if c.IsAjax() {
				c.ApiJSON(500, err.Error(), errData)
				return
			}
			c.HTML(500, err.Error()+"<br><br>"+strings.Join(errData, "<br><br>"))
			c.Log.Error("panic", err, errData)
			return
		}
		if c.IsAjax() {
			c.ApiJSON(500, "Panic", nil)
			return
		}
		c.String(500, "Panic")
	}))
	// r.Use(demoMiddleware())

	// limiterHandle := limiter.New(10000, func(c *znet.Context) {
	// 	c.String(504, "超过限制")
	// })
	// r.Use(limiterHandle)

	// 注册记录器
	if r.IsDebug() {
		r.Use(inspector(r, "/_inspector"))
	}

	// 最长超时时间
	// r.Use(timeout.New(60 * time.Second))
}

//noinspection GoUnusedFunction
func demoMiddleware() func(c *znet.Context) {
	return func(c *znet.Context) {
		c.Next()
	}
}
