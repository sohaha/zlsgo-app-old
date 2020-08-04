package initialize

import (
	"app/compose"
	"app/service/task"
	"app/web/router"
	"github.com/sohaha/zlsgo/zfile"
)

func InitEngine() {
	// 初始化组合
	compose.Init()

	// 初始化定时任务
	task.Init()

	// 初始化 Web 服务
	router.Init()
}

func Clear() {
	// 移除生成的配置文件
	zfile.Rmdir("conf.yml")
}
