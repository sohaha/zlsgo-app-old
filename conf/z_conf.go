package conf

import (
	"fmt"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sohaha/gconf"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
)

type (
	base struct {
		Name  string `mapstructure:"project"`
		Debug bool   // 开启调试模式
		Watch bool   // 监听配置文件变化
	}

	web struct {
		Port       string // 项目端口
		Tls        bool   // 开启 https
		TlsPort    string // https 端口
		Key        string // 证书
		Cert       string // 证书
		Debug      bool   // 开启调试模式
		Pprof      bool   // 开启pprof
		PprofToken string // PprofToken
	}

	// 数据库配置
	db struct {
		Debug              bool
		DBType             string
		MaxLifetime        int
		MaxOpenConns       int
		MaxIdleConns       int
		Prefix             string
		DisableAutoMigrate bool
		Drive              struct {
			MySQL    mysql
			Postgres postgres
			Sqlite3  sqlite3
		}
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
	// sqlite3 sqlite3配置参数
	sqlite3 struct {
		Path string
	}
)

var (
	Log          = zlog.New()
	data         config
	cfg          *gconf.Confhub
	onec         sync.Once
	reloadConfFn []func()
	EnvPort      = ""
	EnvDebug     = false
)

func Init() {
	onec.Do(func() {
		initConf()
		setDebugMode()
		setWatchConf()
		setLogger()
	})
}

func initConf() {
	cfg = gconf.New("conf.yml")
	setDefaultConf(cfg)
	err := cfg.Read()
	zutil.CheckErr(err, true)
	err = cfg.Unmarshal(&data)
	zutil.CheckErr(err, true)
}

// 设置开发模式
func setDebugMode() {
	if !data.Base.Debug && EnvDebug {
		data.Base.Debug = true
	}
}

func setWatchConf() {
	if data.Base.Watch {
		cfg.ConfigChange(func(e fsnotify.Event) {
			_ = cfg.Unmarshal(&data)
			if len(reloadConfFn) > 0 {
				for i := range reloadConfFn {
					go reloadConfFn[i]()
				}
			}
		})
	}
}

func setLogger() {
	Log.ResetFlags(zlog.BitTime | zlog.BitLevel)
	if data.Base.Debug {
		Log.SetLogLevel(zlog.LogDump)
	} else {
		Log.SetLogLevel(zlog.LogSuccess)
	}
}

func SetReloadConf(fn func()) {
	reloadConfFn = append(reloadConfFn, fn)
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
func (a sqlite3) DSN() string {
	return zfile.RealPath(a.Path)
}
