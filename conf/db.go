package conf

import (
	"fmt"

	"github.com/sohaha/zlsgo/zfile"
)

type (
	DBConf struct {
		Debug              bool
		Type               string `mapstructure:"type"`
		MaxLifetime        int
		MaxOpenConns       int
		MaxIdleConns       int
		Prefix             string
		DisableAutoMigrate bool
		MySQL              mysql
		Postgres           postgres
		Sqlite             sqlite
	}
	// mysql mysql 配置参数
	mysql struct {
		Host       string
		Port       int
		User       string
		Password   string
		DBName     string
		Parameters string
	}
	// postgres postgres 配置参数
	postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	// sqlite sqlite 配置参数
	sqlite struct {
		Path string
	}
)

func init() {
	allConf["db"] = &DBConf{
		Type:   "none",
		Debug:  false,
		Prefix: "z_",
		MySQL: mysql{
			Host:       "127.0.0.1",
			Port:       3306,
			User:       "root",
			Password:   "666666",
			DBName:     "zls",
			Parameters: "charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
}

// DB DB配置
func DB() DBConf {
	conf := getConf(func() interface{} {
		conf, ok := allConf["db"]
		if !ok {
			return DBConf{}
		}
		return *conf.(*DBConf)
	})

	return conf.(DBConf)
}

// DSN 数据库连接串
func (a mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

// DSN 数据库连接串
func (a postgres) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		a.Host, a.Port, a.User, a.DBName, a.Password, a.SSLMode)
}

// DSN 数据库连接串
func (a sqlite) DSN() string {
	return zfile.RealPath(a.Path)
}
