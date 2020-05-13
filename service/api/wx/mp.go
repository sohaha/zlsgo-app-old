package wx

import (
	"strconv"
	"time"

	"app/common"
	"github.com/sohaha/wechat"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
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

func (m *Mp) ReceiveMessage(c *znet.Context) {
	body, _ := c.GetDataRaw()
	reply, err := common.Wx.Reply(c.GetAllQuerystMaps(),
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
	replyXml := received.ReplyCustom(func(r *wechat.ReplySt) (xml string) {
		xml, _ = wechat.FormatMap2XML(map[string]string{
			"Content":      "收到:" + r.MsgType + "|" + r.Content,
			"CreateTime":   strconv.FormatInt(time.Now().Unix(), 10),
			"ToUserName":   r.FromUserName,
			"FromUserName": r.ToUserName,
			"MsgType":      "text",
		})
		return
	})

	c.String(200, replyXml)
}
