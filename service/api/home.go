package api

import (
	"os"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztime"
)

// Home Home
type Home struct{}

func (*Home) Home(c *znet.Context) {
	c.ApiJSON(200, "服务正常", map[string]interface{}{
		"time": ztime.Now(),
		"pid":  os.Getpid(),
	})
}

func (*Home) Token(c *znet.Context) {
	panic(333)
}
