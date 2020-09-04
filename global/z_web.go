package global

import (
	"github.com/sohaha/gconf"
)

type (
	stWebConf struct {
		Port       string // 项目端口
		Tls        bool   // 开启 https
		TlsPort    string // https 端口
		Key        string // 证书
		Cert       string // 证书
		Debug      bool   // 开启调试模式
		Pprof      bool   // 开启 pprof
		PprofToken string // pprof Token
	}
)

func (*stWebConf) ConfName(key ...string) string {
	if len(key) > 0 {
		return "web." + key[0]
	}
	return "web"
}

var webConf stWebConf

func (*stCompose) WebDefaultConf(cfg *gconf.Confhub) {
	// web 配置
	for k, v := range map[string]interface{}{
		"port": "3788",
		// "debug": true,
		// "pprof": false,
		// "pprofToken": "",
	} {
		cfg.SetDefault(webConf.ConfName()+"."+k, v)
	}
}

func (*stCompose) WebReadConf(cfg *gconf.Confhub) error {
	// 默认 web debug 开启
	webConf.Debug = true
	err := cfg.Core.UnmarshalKey(webConf.ConfName(), &webConf)
	return err
}

// noinspection GoExportedFuncWithUnexportedType
func WebConf() stWebConf {
	confLock.RLock()
	defer confLock.RUnlock()
	return webConf
}
