package module

import (
	"github.com/sohaha/gconf"
)

type (
	stDemoConf struct{}
)

func (*stDemoConf) ConfName() string {
	return "demo"
}

var demoConf stDemoConf

func (*stCompose) DemoDefaultConf(cfg *gconf.Confhub) {
	// 设置生成的默认配置
	// cfg.SetDefault("demo.description", "is demo description")
}

func (*stCompose) DemoReadConf(cfg *gconf.Confhub) error {
	return cfg.Core.UnmarshalKey(demoConf.ConfName(), &demoConf)
}

func (*stCompose) DemoDone() error {
	return nil
}
