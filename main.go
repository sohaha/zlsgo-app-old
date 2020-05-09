package main

import (
	"app/common"
	"app/conf"
	"app/service"
	"app/service/router"
	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zutil"
)

var (
	port  = zcli.SetVar("port", "Web port").String()
	debug = zcli.SetVar("debug", "Debug mode").Bool()
)

func main() {
	// 设置应用信息
	zcli.Logo = `
 ____ __   ____   __  ____ ____
(__  |  ) / ___) / _\(  _ (  _ \
/ _// (_/\___ \/    \) __/) __/
(____)____(____/\_/\_(__) (__) `
	zcli.Version = "1.0.0"
	// run()
	err := zcli.LaunchServiceRun("ZlsApp", "", run)

	zutil.CheckErr(err, true)
}

func run() {
	// 设置终端执行参数
	conf.EnvDebug = *debug
	conf.EnvPort = *port

	// 初始化
	service.InitEngine()

	// 设置 Web 关闭后回收操作
	router.Closed(stop)

	// 启动 Web 服务
	router.Run()
}

func stop() {
	if saveErr := common.SaveWxCacheData(); saveErr != nil {
		common.Log.Error(saveErr)
	}
}
