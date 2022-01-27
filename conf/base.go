package conf

import (
	"reflect"
	"sync"

	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zutil"
)

type IConf interface {
}

// FileName 配置文件名
const FileName = "conf.yml"

// LogPrefix 日志前缀
const LogPrefix = "[App] "

// noinspection GoUnusedGlobalVariable
var (
	cfg      *gconf.Confhub
	lock     sync.RWMutex
	onec     sync.Once
	EnvPort  = ""
	EnvDebug = false
	allConf  = make(map[string]interface{})
)

func Init(filename string) {
	onec.Do(func() {
		if filename == "" {
			filename = FileName
		}
		cfg = gconf.New(filename)
		for key, conf := range allConf {
			confMaps := make(map[string]interface{})
			typeOf := reflect.TypeOf(conf).Elem()
			if typeOf.Kind() != reflect.Struct {
				continue
			}
			valueOf := reflect.ValueOf(conf).Elem()
			for i := 0; i < typeOf.NumField(); i++ {
				value := valueOf.Field(i)
				field := typeOf.Field(i)
				if value.IsZero() {
					continue
				}
				confMaps[field.Name] = valueOf.Field(i).Interface()
			}
			cfg.SetDefault(key, confMaps)
		}
		zutil.CheckErr(cfg.Read(), true)
		// fix: viper default config invalid
		for key, v := range cfg.GetAll() {
			conf, _ := allConf[key]
			_ = cfg.Core.UnmarshalKey(key, conf)
			cfg.Core.Set(key, v)
		}
		setDebugMode()
	})
}

func setConf(fn func()) {
	lock.Lock()
	defer lock.Unlock()
	fn()
}

func getConf(fn func() interface{}) interface{} {
	lock.RLock()
	defer lock.RUnlock()
	return fn()
}

type BaseConf struct {
	Name        string `mapstructure:"name"`         // 项目名称
	Debug       bool   `mapstructure:"debug"`        // 开启全局调试模式
	LogDir      string `mapstructure:"log_dir"`      // 日志目录
	LogPosition bool   `mapstructure:"log_position"` // 调试下打印日志显示输出位置
}

func init() {
	allConf["base"] = &BaseConf{
		Name: "ZlsApp",
	}
}

// Base Base配置
func Base() BaseConf {
	conf := getConf(func() interface{} {
		conf, ok := allConf["base"]
		if !ok {
			return BaseConf{}
		}
		return *conf.(*BaseConf)
	})

	return conf.(BaseConf)
}

// 设置开发模式
func setDebugMode() {
	base := Base()
	if !base.Debug && EnvDebug {
		setConf(func() {
			base.Debug = true
			allConf["base"] = base
		})
	}
}
