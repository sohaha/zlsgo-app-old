package initialize

import (
	"github.com/sohaha/zstatic"

	"app/global"
	"github.com/zlsgo/resource"
)

func InitManage() {
	c := global.ManageConf()
	if _, e := zstatic.MustBytes(c.Path + "index.html"); e == nil {
		if global.BaseConf().Debug {
			global.Log.Tipsf("TmpFile: %sindex.html already exists\n", c.Path)
		}
		return
	}
	var err error
	defer func() {
		if err != nil {
			global.Log.Error("Manage init failed")
		} else {
			global.Log.Success("Manage init complete")
		}
	}()
	r := resource.New(c.Remote)
	r.SetMd5(c.Md5)
	r.SetDeCompressPath(c.Path)
	r.SetFilterRule([]string{"(.*)/\\.git/", "(.*)/\\.vscode/", "(.*)/\\.idea/"})
	err = r.SilentRun(func(current, total int64) {
	})
}
