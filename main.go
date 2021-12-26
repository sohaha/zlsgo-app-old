package main

import (
	"github.com/sohaha/zlsgo/znet"
)

func main() {
	// 获取一个实例
	r := znet.New()

	// 设置为开发模式
	r.SetMode(znet.DebugMode)

	// 异常处理
	r.PanicHandler(func(c *znet.Context, err error) {
		e := err.Error()
		c.String(500, e)
	})

	// 注册路由
	r.GET("/json", func(c *znet.Context) {
		c.JSON(200, znet.Data{"message": "Hello World"})
	})

	r.GET("/", func(c *znet.Context) {
		c.String(200, "Hello world")
	})

	// 启动
	znet.Run()
}
