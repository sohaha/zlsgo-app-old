package conf

import (
	"github.com/sohaha/gconf"
)

type (
	// 配置文件结构体
	config struct {
		Base     base
		Database db
		Web      web
	}
)

func Base() base {
	return data.Base
}

func Db() db {
	return data.Database
}

func Web() web {
	return data.Web
}

//noinspection GoUnusedExportedFunction
func Data() config {
	return data
}

// 设置初始化配置
func setDefaultConf(cfg *gconf.Confhub) {
	// 基础配置
	cfg.SetDefault("base.name", "App")
	cfg.SetDefault("base.debug", false)

	// web 配置
	cfg.SetDefault("web.port", "3788")
	cfg.SetDefault("web.debug", true)

	// 数据库配置
	cfg.SetDefault("database.prefix", "z_")
	cfg.SetDefault("database.debug", false)

	cfg.SetDefault("database.dbtype", "none")
	cfg.SetDefault("database.drive.sqlite3.path", "./db.sqlite3")

	cfg.SetDefault("database.drive.mysql.host", "127.0.0.1")
	cfg.SetDefault("database.drive.mysql.port", "3306")
	cfg.SetDefault("database.drive.mysql.dbname", "dbname")
	cfg.SetDefault("database.drive.mysql.user", "root")
	cfg.SetDefault("database.drive.mysql.password", "")
	cfg.SetDefault("database.drive.mysql.parameters", "charset=utf8mb4")
}
