package global

import (
	"app/conf"
	"github.com/sohaha/zlsgo/zlog"
)

var (
	Log = zlog.New(conf.LogPrefix)
)

func init() {
	Log.ResetFlags(zlog.BitLevel | zlog.BitTime)

	initMaps = append(initMaps, func() error {
		zlog.Log = Log

		return nil
	})
}
