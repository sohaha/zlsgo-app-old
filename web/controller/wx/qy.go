package wx

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/zlsgo/wechat"

	"app/global"
)

type Qy struct {
}

func (*Qy) GetAccessToken(c *znet.Context) {
	token, err := global.WxQy.GetAccessToken()
	if err != nil {
		c.ApiJSON(211, wechat.ErrorMsg(err), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"time":        global.WxMp.GetAccessTokenExpiresInCountdown(),
		"accessToken": token,
	})
}

func (*Qy) GetJsapiTicket(c *znet.Context) {
	jsapiTicket, err := global.WxQy.GetJsapiTicket()
	if err != nil {
		c.ApiJSON(211, wechat.ErrorMsg(err), nil)
		return
	}
	url := c.Host(true)
	jsSign, err := global.WxQy.GetJsSign(url)
	if err != nil {
		c.ApiJSON(211, wechat.ErrorMsg(err), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"jsapiTicket": jsapiTicket,
		"jsSign":      jsSign,
		"url":         url,
	})
}

func (*Qy) AnyReceiveMessage(c *znet.Context) {
	body, _ := c.GetDataRaw()
	reply, err := global.WxQy.Reply(c.GetAllQuerystMaps(),
		zstring.String2Bytes(body))
	if err != nil {
		c.String(211, err.Error())
		return
	}
	if c.Request.Method == "GET" {
		// Get 请求是响应微信发送的Token验证
		validMsg, err := reply.Valid()
		if err != nil {
			c.String(211, err.Error())
			return
		}
		c.String(200, validMsg)
		return
	}
	received, err := reply.Data()
	if err != nil {
		c.String(211, err.Error())
		return
	}
	c.Log.Info(received)
	replyXml := received.ReplyText("收到消息: " + received.MsgType)

	c.String(200, replyXml)
}
