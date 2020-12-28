package command

import (
	"errors"

	"app/global"
	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zfile"
)

type InitCli struct {
	Force *bool
}

func (i *InitCli) Flags(*zcli.Subcommand) {
	i.Force = zcli.SetVar("force", "覆盖原配置文件").Bool()
}

func (i *InitCli) Run([]string) {
	var err error

	defer func() {
		if err != nil {
			global.Log.Error(err)
		}
	}()

	if zfile.FileExist(global.FileName) {
		if !*i.Force {
			global.Log.Warn("如需覆盖原配置请使用 --force")
			err = global.SaveConf()
			global.Log.Success("配置文件更新成功")
			return
		}
		zfile.Rmdir(global.FileName)
	}
	// 配置初始化
	global.ReadConf(false)
	if !zfile.FileExist(global.FileName) {
		err = errors.New("配置文件生成失败")
	} else {
		global.Log.Success("配置文件生成成功")
	}
}
