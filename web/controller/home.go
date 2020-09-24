package controller

import (
	"os"

	"app/logic"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztime"
)

// Home Home 控制器
type Home struct{}

func (*Home) Get(c *znet.Context) {
	c.ApiJSON(200, "服务正常", map[string]interface{}{
		"time": ztime.Now(),
		"pid":  os.Getpid(),
	})
}

func (*Home) GetServerInfo(c *znet.Context) {
	if !c.Engine.IsDebug() {
		c.Log.Warn("非调试模式默认无法查看服务器信息")
		c.ApiJSON(403, "没有权限", nil)
	}

	c.ApiJSON(200, "服务器信息", logic.GetServerInfo())
}
