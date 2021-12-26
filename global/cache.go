package global

import (
	"github.com/sohaha/zlsgo/zcache"
)

var (
	Cache *zcache.Table
)

func init() {
	initMaps = append(initMaps, func() error {
		Cache = zcache.New("app")
		return nil
	})
}
