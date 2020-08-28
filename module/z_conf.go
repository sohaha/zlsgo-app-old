package module

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/sohaha/gconf"

	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
)

type (
	stCompose  struct{}
	stBaseConf struct {
		Name        string `mapstructure:"project"`
		Debug       bool   `mapstructure:"debug"`        // 开启调试模式
		Watch       bool   `mapstructure:"watch"`        // 监听配置文件变化
		LogDir      string `mapstructure:"log_dir"`      // 日志目录
		LogPosition bool   `mapstructure:"log_position"` // 调试下打印日志显示输出位置
	}
)

func (b *stBaseConf) ConfName() string {
	return "base"
}

const FileName = "conf.yml"

// noinspection GoUnusedGlobalVariable
var (
	baseConf stBaseConf
	cfg      *gconf.Confhub
	onec     sync.Once
	onecInit sync.Once
	confLock sync.RWMutex
	Log      = zlog.New("[App] ")
	Cache    = zcache.New("app")
	EnvPort  = ""
	EnvDebug = false
)

func init() {
	Log.ResetFlags(zlog.BitLevel | zlog.BitTime)
}

func Init() {
	Read(true)
	onec.Do(func() {
		setDebugMode()
		// setWatchConf()
	})
}

func Read(init bool) {
	onecInit.Do(func() {
		zutil.Try(func() {
			cfg = gconf.New(FileName)
			setComposeDefaultConf()
			readComposeConf()
			if init {
				initCompose()
			}
			setLogger()
		}, func(e interface{}) {
			if err, ok := e.(error); ok {
				Log.Fatal(err.Error())
			}
		})
	})
}

// 设置初始化模块
func setComposeDefaultConf() {
	err := zutil.RunAssignMethod(&stCompose{}, func(methodName string) bool {
		return strings.HasSuffix(methodName, "DefaultConf")
	}, cfg)
	zutil.CheckErr(err)
}

// 读取模块配置
func readComposeConf() {
	err := cfg.Read()
	zutil.CheckErr(err)
	confLock.Lock()
	defer confLock.Unlock()
	// fix: viper default config invalid
	for key, v := range cfg.GetAll() {
		cfg.Core.Set(key, v)
	}
	err = zutil.RunAssignMethod(&stCompose{}, func(methodName string) bool {
		return strings.HasSuffix(methodName, "ReadConf")
	}, cfg)
	zutil.CheckErr(err)
	// Update the current configuration to the configuration file
	err = cfg.Core.WriteConfig()
	if err != nil {
		Log.Warn(err)
	}
}

// 模块配置
func initCompose() {
	err := zutil.RunAssignMethod(&stCompose{}, func(methodName string) bool {
		return strings.HasSuffix(methodName, "Done")
	})
	zutil.CheckErr(err)
}

// 设置开发模式
func setDebugMode() {
	if !BaseConf().Debug && EnvDebug {
		SetConfData(func() {
			baseConf.Debug = true
		})
	}
}

// func setWatchConf() {
// 	if BaseConf().Watch {
// 		cfg.ConfigChange(func(e fsnotify.Event) {
// 			data.Lock()
// 			_ = cfg.Unmarshal(&data)
// 			data.Unlock()
// 			if len(reloadConfFn) > 0 {
// 				for i := range reloadConfFn {
// 					go reloadConfFn[i]()
// 				}
// 			}
// 		})
// 	}
// }

func setLogger() {
	if BaseConf().Debug {
		Log.SetLogLevel(zlog.LogDump)
		flags := zlog.BitTime | zlog.BitLevel
		if baseConf.LogPosition {
			flags = flags | zlog.BitLongFile
		}
		Log.ResetFlags(flags)
	} else {
		Log.SetLogLevel(zlog.LogSuccess)
		Log.ResetFlags(zlog.BitTime | zlog.BitLevel)
	}
	if BaseConf().LogDir != "" {
		Log.SetSaveFile(filepath.Join(BaseConf().LogDir, "app.log"), true)
	}
}

func SetConfData(fn func()) {
	confLock.Lock()
	defer confLock.Unlock()
	fn()
}

// func SetReloadConf(fn func()) {
// 	reloadConfFn = append(reloadConfFn, fn)
// }

func (*stCompose) BaseDefaultConf(cfg *gconf.Confhub) {
	// 基础配置
	cfg.SetDefault(baseConf.ConfName(), map[string]interface{}{
		"debug":        false,
		"log_dir":      "",
		"log_position": true,
	})
}

// noinspection GoExportedFuncWithUnexportedType
func BaseConf() stBaseConf {
	confLock.RLock()
	defer confLock.RUnlock()
	return baseConf
}

func (*stCompose) BaseReadConf(cfg *gconf.Confhub) error {
	return cfg.Core.UnmarshalKey(baseConf.ConfName(), &baseConf)
}

// GetConfAll 获取配置值
// noinspection ALL
func GetConfAll() map[string]interface{} {
	confLock.RLock()
	defer confLock.RUnlock()
	return cfg.GetAll()
}

// GetConfInstance 获取配置实例
// noinspection ALL
func GetConfInstance() *gconf.Confhub {
	return cfg
}
