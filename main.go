package main

import (
	"app/module"
	"app/module/initialize"
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
 ____ __   ____   __  ____ ____
(__  |  ) / ___) / _\(  _ (  _ \
/ _// (_/\___ \/    \) __/) __/
(____)____(____/\_/\_(__) (__) `
	zcli.Version = "1.0.0"
	zcli.Lang = "zh"
	zcli.Add("init", "生成配置", &InitCli{})

	err := zcli.LaunchServiceRun(zcli.Name, "", run)

	zutil.CheckErr(err, true)
}

func run() {
	// 设置终端执行参数
	module.EnvDebug = *debug
	module.EnvPort = *port

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
	if zfile.FileExist(module.FileName) {
		if !*i.Force {
			module.Log.Warn("配置文件已存在，如需覆盖原配置请使用 --force")
			return
		}
		zfile.Rmdir(module.FileName)
	}
	// 配置初始化
	module.Read(false)
	if zfile.FileExist(module.FileName) {
		module.Log.Success("配置文件初始化成功")
	} else {
		module.Log.Error("配置文件初始化失败")
	}
}
