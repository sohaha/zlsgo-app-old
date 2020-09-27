package initialize

import (
	"github.com/sohaha/zlsgo/zfile"

	"app/global"
	"app/global/task"
	"app/web/router"
)

// InitEngine 初始化模块
func InitEngine() {
	// 初始化组合
	global.InitConf()

	// 初始化定时任务
	task.Init()

	// 初始化 Web 服务
	router.Init()
}

// Clear 移除生成的配置文件
func Clear() {
	zfile.Rmdir("conf.yml")
}
