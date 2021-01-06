package wx

import (
	"strings"
	"sync"

	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"

	"github.com/zlsgo/wechat"

	"app/global"
)

type Open struct {
}

func (*Open) AnyNotification(c *znet.Context) {
	data, _ := c.GetDataRaw()
	_, err := global.WxOpen.ComponentVerifyTicket(data)
	if err != nil {
		c.Log.Warn(err.Error())
	}
	// 需要返回 success 给微信
	c.String(200, "success")
}

// 公众号授权
func (*Open) GetApiQueryAuth(c *znet.Context) {
	authCode, ok := c.GetQuery("auth_code")
	if !ok {
		c.Log.Debug("需要跳转")
	}
	redirectUri := c.Host(true)
	// todo 这里需要把 http://mac.zj.73zls.com 换成你项目的域名
	redirectUri = strings.Replace(redirectUri, c.Host(), "http://mac.zj.73zls.com", 1)
	res, redirect, err := global.WxOpen.ComponentApiQueryAuth(authCode, redirectUri)
	if err != nil {
		if err != wechat.ErrOpenJumpAuthorization {
			c.ApiJSON(211, wechat.ErrorMsg(err), nil)
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

func (*Open) GetTicket(c *znet.Context) {
	ticket, err := global.WxOpen.GetConfig().(*wechat.Open).GetComponentTicket()
	c.Log.Debug(ticket, err)
	if err != nil {
		c.ApiJSON(211, wechat.ErrorMsg(err), nil)
		return
	}
	c.ApiJSON(200, "获取 Ticket", ticket)
}

var o sync.Once

func (*Open) GetAccessToken(c *znet.Context) {
	o.Do(func() {
		// todo 因为开放平台的是授权的时候获取的，这里可以强制手动设置
		// global.WxOpen.GetConfig().(*wechat.Open).SetAuthorizerAccessToken("", "", "", 4)
	})
	token, err := global.WxOpen.GetAccessToken()
	if err != nil {
		c.ApiJSON(211, wechat.ErrorMsg(err), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"time":        global.WxOpen.GetAccessTokenExpiresInCountdown(),
		"accessToken": token,
	})
}

func (*Open) GetJsapiTicket(c *znet.Context) {
	jsapiTicket, err := global.WxOpen.GetJsapiTicket()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	url := c.Host(true)
	jsSign, err := global.WxOpen.GetJsSign(url)
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
