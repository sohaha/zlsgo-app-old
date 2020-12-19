package initialize

import (
	"app/global"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zhttp"
	"github.com/sohaha/zlsgo/zutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ing(format string, v ...interface{}) {
	global.Log.Warnf(format+"\n", v...)
}

func InitManage() {
	c := global.ManageConf()
	p := zfile.RealPath(c.Path)
	if zfile.DirExist(p) {
		return
	}
	url := c.Gz
	name := path.Base(url)
	ext := path.Ext(name)
	ext = zutil.IfVal(ext == "", "zip", strings.TrimLeft(ext, ".")).(string)
	tmp := zfile.TmpPath()
	tmpFile := tmp + "/" + name
	global.Log.Error("需要初始化后台", c, p, tmp, name, ext)
	ing("开始初始化后台")
	res, err := zhttp.Get(url, zhttp.DownloadProgress(func(current, total int64) {
	}))
	if err != nil {
		global.Log.Error(err)
		return
	}
	err = res.ToFile(tmpFile)
	if err != nil {
		global.Log.Error(err)
		return
	}
	tmp = tmp + "/tmp"
	err = zfile.GzDeCompress(tmpFile, tmp)
	if err != nil {
		global.Log.Errorf("文件解压失败: %s\n", err)
		return
	}
	cf := tmp + "/init-config.json"
	if zfile.FileExist(cf) {

	}
	filepath.Walk(tmp, func(path string, i os.FileInfo, err error) error {
		dest := strings.Replace(path, tmp, p, 1)
		if i.IsDir() {
			os.MkdirAll(dest, i.Mode())
			return nil
		}
		_ = zfile.CopyFile(path, dest)
		return nil
	})
	global.Log.Success("后台初始化完成")
}
