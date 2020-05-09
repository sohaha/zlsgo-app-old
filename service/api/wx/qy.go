package wx

import (
	"app/common"
	"github.com/sohaha/zlsgo/znet"
)

type Qy struct {
}

func (*Qy) AccessToken(c *znet.Context) {
	token, err := common.WxQy.GetAccessToken()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"time":        common.Wx.GetAccessTokenExpiresInCountdown(),
		"accessToken": token,
	})
}

func (*Qy) JsapiTicket(c *znet.Context) {
	jsapiTicket, err := common.WxQy.GetJsapiTicket()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	url := c.Host(true)
	jsSign, err := common.WxQy.GetJsSign(url)
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
