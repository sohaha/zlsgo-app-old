package router

import (
	"github.com/sohaha/zlsgo/znet"
	_ "github.com/sohaha/zlsgo/znet/limiter"
)

func registerMiddleware(_ *znet.Engine) {
	// r.Use(demoMiddleware())

	// limiterHandle := limiter.New(10000, func(c *znet.Context) {
	// 	c.String(504, "超过限制")
	// })
	// r.Use(limiterHandle)
}

//noinspection GoUnusedFunction
func demoMiddleware() func(c *znet.Context) {
	return func(c *znet.Context) {
		c.Next()
	}
}
