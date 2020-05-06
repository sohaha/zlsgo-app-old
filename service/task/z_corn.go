package task

import (
	"reflect"

	"github.com/sohaha/zlsgo/ztime/cron"
)

type cornTask struct {
}

func Init() {
	object := reflect.ValueOf(&cornTask{})
	taskLen := object.NumMethod()
	if taskLen == 0 {
		return
	}
	corntabObj := cron.New()
	for i := 0; i < taskLen; i++ {
		v := object.Method(i)
		v.Call([]reflect.Value{
			reflect.ValueOf(corntabObj),
		})
	}
	corntabObj.Run()
}
