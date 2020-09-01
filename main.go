package main

import (
	"app/global"
	"app/global/initialize"
	"app/web/router"

	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zfile"
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
	zcli.SetLangText("zh", "init", "生成配置")
	zcli.Add("init", zcli.GetLangText("init", "Init config file"), &InitCli{})

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

type InitCli struct {
	Force *bool
}

func (i *InitCli) Flags(*zcli.Subcommand) {
	i.Force = zcli.SetVar("force", "覆盖原配置文件").Bool()
}

func (i *InitCli) Run([]string) {
	if zfile.FileExist(global.FileName) {
		if !*i.Force {
			global.Log.Warn("配置文件已存在，如需覆盖原配置请使用 --force")
			return
		}
		zfile.Rmdir(global.FileName)
	}
	// 配置初始化
	global.Read(false)
	if zfile.FileExist(global.FileName) {
		global.Log.Success("配置文件初始化成功")
	} else {
		global.Log.Error("配置文件初始化失败")
	}
}
