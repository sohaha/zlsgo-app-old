package wx

import (
	"strings"
	"sync"

	"app/common"
	"github.com/sohaha/wechat"
	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
)

type Open struct {
}

func (*Open) Notification(c *znet.Context) {
	data, _ := c.GetDataRaw()
	_, err := common.WxOpen.ComponentVerifyTicket(data)
	if err != nil {
		c.Log.Warn(err.Error())
	}
	c.String(200, "")
}

// 公众号授权
func (*Open) ApiQueryAuth(c *znet.Context) {
	authCode, ok := c.GetQuery("auth_code")
	if !ok {
		c.Log.Debug("需要跳转")
	}
	redirectUri := c.Host(true)
	// todo dev
	redirectUri = strings.Replace(redirectUri, c.Host(),
		"http://mac.zj.73zls.com", 1)
	res, redirect, err := common.WxOpen.ComponentApiQueryAuth(authCode,
		redirectUri)
	if err != nil {
		if err != wechat.ErrOpenJumpAuthorization {
			c.ApiJSON(211, err.Error(), nil)
			return
		}
		// JS 发起跳转
		c.Template(200, `<html lang='zh'><head><title>Loading...
</title></head><body><script type='text/javascript'>
referLink=document.createElement('a');referLink.href="{{.redirect}}";referLink.
click()</script></body></html>`, map[string]string{"redirect": redirect})
		return
	}
	c.ApiJSON(200, "公众号授权成功", zjson.Parse(res).Value())
}

func (*Open) Ticket(c *znet.Context) {
	ticket, err := common.WxOpen.GetConfig().(*wechat.Open).GetComponentTicket()
	c.Log.Debug(ticket, err)
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	c.ApiJSON(200, "获取 Ticket", ticket)
}

var o sync.Once
func (*Open) AccessToken(c *znet.Context) {
	o.Do(func() {
		// 因为开放平台的是授权的时候获取的，需要根据实际情况调用
		common.WxOpen.GetConfig().(*wechat.Open).SetAuthorizerAccessToken(
			"wx2d9176b8947d2103",
			"33_4bYCZQlAYKZ7MHq4P_LmGZEzdsc_RRFtzWjtUv5iPJmJIxsxBbFkMdKuu1lW2WlmlDldQK1aD2NyozO0b8JX--czTHk0BcbNMm8cdR0G2u_mCXftuzH3Ymgls96CaeqmN2BNmVYxhxbYqEP6FSMbAGDPCW", "refreshtoken@@@kuo1_wz2vPtE3dOtNKLbsgk5Hotc4Z3SdRYms0tgEGU", 4)
	})
	token, err := common.WxOpen.GetAccessToken()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"time":        common.WxOpen.GetAccessTokenExpiresInCountdown(),
		"accessToken": token,
	})
}

func (*Open) JsapiTicket(c *znet.Context) {
	jsapiTicket, err := common.WxOpen.GetJsapiTicket()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	url := c.Host(true)
	jsSign, err := common.WxOpen.GetJsSign(url)
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
