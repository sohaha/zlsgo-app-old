package main

import (
	"app/command"
	"app/global"
	"app/global/initialize"
	"app/web/router"

	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zutil"
)

var (
	port  = zcli.SetVar("port", "Web 服务端口").String()
	debug = zcli.SetVar("debug", "开启调试模式").Bool()
)

func main() {
	// 设置应用信息
	zcli.Name = "ZlsApp"
	zcli.Logo = `
   _____                   
  /  _  \  ______  ______  
 /  /_\  \ \____ \ \____ \ 
/    |    \|  |_> >|  |_> >
\____|__  /|   __/ |   __/ 
        \/ |__|    |__|     `
	zcli.Version = "1.0.0"
	zcli.Lang = "zh"

	command.Reg()
	err := zcli.LaunchServiceRun(zcli.Name, "", run)
	zutil.CheckErr(err, true)
}

func run() {
	// 设置终端执行参数
	global.EnvDebug = *debug
	global.EnvPort = *port
	// 初始化
	initialize.InitEngine()

	// 启动 Web 服务
	router.Run()
}
