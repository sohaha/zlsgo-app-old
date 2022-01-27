package global

import (
	"path/filepath"

	"app/conf"
	"github.com/sohaha/zlsgo/zlog"
)

var (
	Log = zlog.New(conf.LogPrefix)
)

func init() {
	Log.ResetFlags(zlog.BitLevel)

	initMaps = append(initMaps, func() error {
		zlog.Log = Log
		baseConf := conf.Base()
		if baseConf.Debug {
			Log.SetLogLevel(zlog.LogDump)
			flags := zlog.BitTime | zlog.BitLevel
			if baseConf.LogPosition {
				flags = flags | zlog.BitLongFile
			}
			Log.ResetFlags(flags)
		} else {
			Log.SetLogLevel(zlog.LogSuccess)
			Log.ResetFlags(zlog.BitTime | zlog.BitLevel)
		}

		if baseConf.LogDir != "" {
			Log.SetSaveFile(filepath.Join(baseConf.LogDir, "app.log"), true)
		}

		return nil
	})
}
