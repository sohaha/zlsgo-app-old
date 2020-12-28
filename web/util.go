package web

import (
	"fmt"

	"github.com/sohaha/zlsgo/zjson"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zstring"
)

func addJSON(json []byte, data interface{}, more ...map[string]interface{}) []byte {
	json, _ = zjson.SetBytes(json, "data", data)
	if len(more) > 0 {
		for _, d := range more {
			for k, v := range d {
				json, _ = zjson.SetBytes(json, "data."+k, v)
			}
		}
	}
	return json
}

func ApiJSON(c *znet.Context, code int, msg string, data interface{}, more ...map[string]interface{}) {
	raw := zstring.String2Bytes(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
	json := addJSON(raw, data, more...)
	c.SetContentType(znet.ContentTypeJSON)
	c.Byte(200, json)
	return
}
