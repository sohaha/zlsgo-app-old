package test

import (
	"os"
	"path/filepath"

	"app/global"
	"github.com/sohaha/zlsgo/zfile"
)

func Run(fn func()) {
	fileList := make([]string, 0)

	walkDir(func(path string) {
		fileList = append(fileList, path)
	})

	path := "../conf.yml"
	for i := 0; i < 3; i++ {
		if zfile.FileExist(path) {
			// 如果项目存在配置文件，直接 copy 一份
			_ = zfile.CopyFile(path, "./conf.yml")
			break
		}
		path = "../" + path
	}

	// 清除数据
	defer func() {
		walkDir(func(path string) {
			rm := path == "conf.yml"
			if !rm {
				for i, v := range fileList {
					if path == v {
						rm = false
						fileList = append(fileList[:i], fileList[i+1:]...)
						break
					}
				}
			}
			if rm {
				zfile.Rmdir(path)
			}
		})
	}()

	// 初始化组合
	global.InitConf()

	// 开始执行
	fn()
}

func walkDir(f func(path string)) {
	_ = filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		f(path)
		return nil
	})
}
