package wx

import (
	"app/common"
	"github.com/sohaha/zlsgo/znet"
)

type Mp struct {
}

func (*Mp) AccessToken(c *znet.Context) {
	token, err := common.Wx.GetAccessToken()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"time":        common.Wx.GetAccessTokenExpiresInCountdown(),
		"accessToken": token,
	})
}

func (*Mp) JsapiTicket(c *znet.Context) {
	jsapiTicket, err := common.Wx.GetJsapiTicket()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	url := c.Host(true)
	jsSign, err := common.Wx.GetJsSign(url)
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"jsapiTicket": jsapiTicket,
		"jsSign":      jsSign,
		"url":         url,
	})
}
