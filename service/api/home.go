package api

import (
	"os"

	"app/common"
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
	token, err := common.Wx.GetAccessToken()
	if err != nil {
		c.ApiJSON(200, err.Error(), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"time":        common.Wx.GetAccessTokenExpiresInCountdown(),
		"accessToken": token,
	})
}
