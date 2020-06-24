package conf

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/sohaha/gconf"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
)

type (
	base struct {
		Name   string `mapstructure:"project"`
		Debug  bool   // 开启调试模式
		Watch  bool   // 监听配置文件变化
		Logdir string // 日志目录
	}

	web struct {
		Port       string // 项目端口
		Tls        bool   // 开启 https
		TlsPort    string // https 端口
		Key        string // 证书
		Cert       string // 证书
		Debug      bool   // 开启调试模式
		Pprof      bool   // 开启 pprof
		PprofToken string // pprof Token
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

const FileName = "conf.yml"

var (
	Log          = zlog.New("[App] ")
	data         config
	cfg          *gconf.Confhub
	onec         sync.Once
	onecInit     sync.Once
	reloadConfFn []func()
	EnvPort      = ""
	EnvDebug     = false
)

func Init() {
	Read()
	onec.Do(func() {
		setDebugMode()
		setWatchConf()
	})
}

func Read() {
	onecInit.Do(func() {
		cfg = gconf.New(FileName)
		setDefaultConf(cfg)
		err := cfg.Read()
		zutil.CheckErr(err, true)
		err = cfg.Unmarshal(&data.value)
		zutil.CheckErr(err, true)
		setLogger()
	})
}

// 设置开发模式
func setDebugMode() {
	if !Base().Debug && EnvDebug {
		SetConfData(func(data *ConfigValue) {
			data.Base.Debug = true
		})
	}
}

func setWatchConf() {
	if Base().Watch {
		cfg.ConfigChange(func(e fsnotify.Event) {
			data.Lock()
			_ = cfg.Unmarshal(&data)
			data.Unlock()
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
	if Base().Debug {
		Log.SetLogLevel(zlog.LogDump)
	} else {
		Log.SetLogLevel(zlog.LogSuccess)
	}
	if Base().Logdir != "" {
		Log.SetSaveFile(filepath.Join(Base().Logdir, "base.log"), true)
	}
}

func SetConfData(fn func(data *ConfigValue)) {
	data.Lock()
	defer data.Unlock()
	fn(data.value)
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
