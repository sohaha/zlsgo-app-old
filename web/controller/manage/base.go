package manage

import (
	"github.com/sohaha/zlsgo/znet"
)

func Authority() func(c *znet.Context) {
	return func(c *znet.Context) {
		c.Next()
	}
}
