package middleware

import (
	"net"
	"strings"

	"github.com/sohaha/zlsgo/znet"

	"app/global"
)

// maintainModeMiddleware 维护模式
//goland:noinspection GoUnusedFunction
func maintainModeMiddleware() func(c *znet.Context) {
	return func(c *znet.Context) {
		ip := c.GetClientIP()
		if ip != "" && global.BaseConf().MaintainMode {
			ipWhitelist := strings.Split(global.BaseConf().IPWhitelist, ",")
			for _, v := range ipWhitelist {
				if v == ip {
					c.Next()
					return
				}
				_, network, err := net.ParseCIDR(v)
				if err == nil && network.Contains(net.ParseIP(ip)) {
					c.Next()
					return
				}
			}
		}
		c.Next()
	}
}
