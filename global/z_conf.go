package global

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/zcli"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
)

type (
	stCompose  struct{}
	stBaseConf struct {
		Name         string `mapstructure:"project"`       // 项目名称
		Debug        bool   `mapstructure:"debug"`         // 开启全局调试模式
		Watch        bool   `mapstructure:"watch"`         // 监听配置文件变化
		LogDir       string `mapstructure:"log_dir"`       // 日志目录
		LogPosition  bool   `mapstructure:"log_position"`  // 调试下打印日志显示输出位置
		MaintainMode bool   `mapstructure:"maintain_mode"` // 维护模式
		IPWhitelist  string `mapstructure:"ip_whitelist"`  // 维护模式下，白名单
	}
)

func (b *stBaseConf) ConfName() string {
	return "base"
}

// FileName 配置文件名
const FileName = "conf.yml"

// 日志前缀
const LogPrefix = "[App] "

// noinspection GoUnusedGlobalVariable
var (
	baseConf stBaseConf
	cfg      *gconf.Confhub
	onec     sync.Once
	onecInit sync.Once
	confLock sync.RWMutex
	Log      = zlog.New(LogPrefix)
	Cache    = zcache.New("app")
	EnvPort  = ""
	EnvDebug = false
)

func init() {
	Log.ResetFlags(zlog.BitLevel | zlog.BitTime)
}

// InitConf 初始化配置
func InitConf() {
	ReadConf(true)
	onec.Do(func() {
		setDebugMode()
		// setWatchConf()
	})
}

// ReadConf 读取配置
func ReadConf(init bool) {
	onecInit.Do(func() {
		zutil.Try(func() {
			cfg = gconf.New(FileName)
			setComposeDefaultConf()
			readComposeConf()
			setLogger()
			if init {
				initCompose()
			}
		}, func(e interface{}) {
			if err, ok := e.(error); ok {
				Log.Fatal(err.Error())
			}
		})
	})
}

// SaveConf 保存当前配置
func SaveConf() error {
	ReadConf(false)
	// Update the current configuration to the configuration file
	return cfg.Core.WriteConfig()
}

// 获取指定配置
//goland:noinspection GoUnusedExportedFunction
func GetConf(key string) interface{} {
	confLock.RLock()
	defer confLock.RUnlock()
	return cfg.Get(key)
}

// 更新配置
//goland:noinspection GoUnusedExportedFunction
func UpdateConf(key string, value interface{}) error {
	confLock.RLock()
	defer confLock.RUnlock()
	cfg.Set(key, value)
	return SaveConf()
}

// GetConfAll 获取配置值
// noinspection ALL
func GetConfAll() map[string]interface{} {
	confLock.RLock()
	defer confLock.RUnlock()
	return cfg.GetAll()
}

// 设置初始化模块
func setComposeDefaultConf() {
	// 读取环境变量 zcliName_xxx
	cfg.Core.SetEnvPrefix(zcli.Name)
	cfg.Core.AutomaticEnv()
	cfg.Core.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 遍历初始化配置
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
		"log_position": false,
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

// GetConfInstance 获取配置实例
// noinspection ALL
func GetConfInstance() *gconf.Confhub {
	return cfg
}

// Recover 回收处理
func Recover() {
	err := zutil.RunAssignMethod(&stCompose{}, func(methodName string) bool {
		return strings.HasSuffix(methodName, "Recover")
	})
	zutil.CheckErr(err)
}
