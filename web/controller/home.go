package controller

import (
	"os"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztime"
)

// Home Home控制器
type Home struct{}

// Get 首页
func (*Home) Get(c *znet.Context) {
	c.ApiJSON(200, "服务正常", map[string]interface{}{
		"time": ztime.Now(),
		"pid":  os.Getpid(),
	})
}
