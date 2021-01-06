package wx

import (
	"strconv"
	"time"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"

	"app/global"
	"github.com/zlsgo/wechat"
)

type Mp struct {
}

func (*Mp) GetAccessToken(c *znet.Context) {
	token, err := global.WxMp.GetAccessToken()
	if err != nil {
		c.Log.Debug(err)
		c.ApiJSON(211, wechat.ErrorMsg(err), nil)
		return
	}
	c.ApiJSON(200, "获取成功", map[string]interface{}{
		"time":        global.WxMp.GetAccessTokenExpiresInCountdown(),
		"accessToken": token,
	})
}

func (*Mp) GetJsapiTicket(c *znet.Context) {
	jsapiTicket, err := global.WxMp.GetJsapiTicket()
	if err != nil {
		c.ApiJSON(211, wechat.ErrorMsg(err), nil)
		return
	}
	url := c.Host(true)
	jsSign, err := global.WxMp.GetJsSign(url)
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

func (m *Mp) AnyReceiveMessage(c *znet.Context) {
	body, _ := c.GetDataRaw()
	reply, err := global.WxMp.Reply(c.GetAllQuerystMaps(),
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
