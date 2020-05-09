package router

import (
	"time"

	"github.com/sohaha/zlsgo/znet"
	_ "github.com/sohaha/zlsgo/znet/limiter"
	"github.com/sohaha/zlsgo/znet/timeout"
)

func registerMiddleware(r *znet.Engine) {
	// r.Use(demoMiddleware())

	// limiterHandle := limiter.New(10000, func(c *znet.Context) {
	// 	c.String(504, "超过限制")
	// })
	// r.Use(limiterHandle)

	// 最长超时时间
	r.Use(timeout.New(60 * time.Second))
}

//noinspection GoUnusedFunction
func demoMiddleware() func(c *znet.Context) {
	return func(c *znet.Context) {
		c.Next()
	}
}
