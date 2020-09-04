package global

import (
	"fmt"
	"strings"

	"github.com/sohaha/gconf"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zvalid"
)

type (
	// 数据库配置
	stDatabaseConf struct {
		Debug              bool
		DBType             string `mapstructure:"db_type"`
		MaxLifetime        int
		MaxOpenConns       int
		MaxIdleConns       int
		Prefix             string
		DisableAutoMigrate bool
		MySQL              mysql
		Postgres           postgres
		Sqlite3            sqlite
	}

	// mysql mysql配置参数
	mysql struct {
		Host       string
		Port       int
		User       string
		Password   string
		DBName     string
		Parameters string
	}
	// postgres postgres配置参数
	postgres struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	// sqlite sqlite配置参数
	sqlite struct {
		Path string
	}
)

func (*stDatabaseConf) ConfName(key ...string) string {
	if len(key) > 0 {
		return "database." + key[0]
	}
	return "database"
}

var (
	databaseConf            stDatabaseConf
	databaseDefaultInitConf = map[string]interface{}{
		"prefix":           "z_",
		"debug":            false,
		"db_type":          "none",
		"mysql.host":       "127.0.0.1",
		"mysql.port":       "3306",
		"mysql.dbname":     "dbname",
		"mysql.user":       "root",
		"mysql.password":   "root",
		"mysql.parameters": "charset=utf8mb4&parseTime=True&loc=Local",
		// "sqlite.path":     "./db.sqlite",
	}
)

func (*stCompose) DatabaseDefaultConf(cfg *gconf.Confhub) {
	// 数据库配置
	for k, v := range databaseDefaultInitConf {
		cfg.SetDefault(databaseConf.ConfName()+"."+k, v)
	}
}

func (*stCompose) DatabaseReadConf(cfg *gconf.Confhub) error {
	return cfg.Core.UnmarshalKey(databaseConf.ConfName(), &databaseConf)
}

// noinspection GoExportedFuncWithUnexportedType
func DatabaseConf() stDatabaseConf {
	confLock.RLock()
	defer confLock.RUnlock()
	return databaseConf
}

func validDb() error {
	// 数据库
	dbTypes := []string{"mysql", "sqlite", "postgres"}
	return zvalid.Text(DatabaseConf().DBType, "数据库类型").Required().
		EnumString(append(dbTypes, "none"), "数据库类型暂只支持: "+strings.Join(dbTypes, ", ")).Error()
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
