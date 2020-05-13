package wx

import (
	"app/common"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
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

func (*Qy) ReceiveMessage(c *znet.Context) {
	body, _ := c.GetDataRaw()
	reply, err := common.WxQy.Reply(c.GetAllQuerystMaps(),
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
