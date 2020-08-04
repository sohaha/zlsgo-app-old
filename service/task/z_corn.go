package task

import (
	"github.com/sohaha/zlsgo/ztime/cron"
	"github.com/sohaha/zlsgo/zutil"
)

type cornTask struct{}

func Init() {
	corntabObj := cron.New()
	err := zutil.RunAllMethod(&cornTask{}, corntabObj)
	zutil.CheckErr(err)
	corntabObj.Run()
}
