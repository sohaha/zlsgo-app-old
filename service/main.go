package service

import (
	"app/conf"
	"app/service/model"
	"app/service/router"
	"app/service/task"
	"github.com/sohaha/zlsgo/zfile"
)

func InitEngine() {
	// 初始化配置
	conf.Init()

	// 初始化数据库
	model.Init()

	// 初始化定时任务
	task.Init()

	// 初始化 Web 服务
	router.Init()
}

func Clear() {
	// 移除生成的配置文件
	zfile.Rmdir("conf.yml")
}
