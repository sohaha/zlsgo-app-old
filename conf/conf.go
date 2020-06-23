/*
 * @Author: seekwe
 * @Date: 2020-06-23 19:16:25
 * @Last Modified by:: seekwe
 * @Last Modified time: 2020-06-23 19:17:32
 */
package conf

import (
	"sync"

	"github.com/sohaha/gconf"
)

type (
	// 配置文件结构体
	config struct {
		sync.RWMutex
		value *ConfigValue
	}
	ConfigValue struct {
		Base     base
		Database db
		Web      web
	}
)

func Base() base {
	data.RLock()
	defer data.RUnlock()
	return data.value.Base
}

func Db() db {
	data.RLock()
	defer data.RUnlock()
	return data.value.Database
}

func Web() web {
	data.RLock()
	defer data.RUnlock()
	return data.value.Web
}

//noinspection GoUnusedExportedFunction
func Data() ConfigValue {
	data.RLock()
	defer data.RUnlock()
	return *data.value
}

// 设置初始化配置
func setDefaultConf(cfg *gconf.Confhub) {
	// 基础配置
	cfg.SetDefault("base.project", "App")
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
