package main

import (
	"app/conf"
	"app/global"
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

	err := zcli.LaunchServiceRun(zcli.Name, "", run)
	zutil.CheckErr(err, true)
}

func run() {
	// 设置终端执行参数
	conf.EnvDebug = *debug
	conf.EnvPort = *port
	conf.Init(conf.FileName)
	global.Init()
	router.Init()
	// logic.All(("app/conf"))
	router.Run()
}
