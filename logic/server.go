package logic

import (
	"runtime"
	
	"github.com/sohaha/zlsgo/zfile"
)

func GetServerInfo() map[string]interface{} {
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH

	return map[string]interface{}{
		"os":     osDic,
	}
}
