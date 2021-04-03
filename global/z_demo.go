package global

import (
	"github.com/zlsgo/conf"
)

type (
	stDemoConf struct{}
)

func (*stDemoConf) ConfName() string {
	return "demo"
}

var demoConf stDemoConf

// 默认配置
func (*stCompose) DemoDefaultConf(cfg *conf.Confhub) {
	// 设置生成的默认配置
	// cfg.SetDefault("demo.description", "is demo description")
}

// 读取配置
func (*stCompose) DemoReadConf(cfg *conf.Confhub) error {
	return cfg.Core.UnmarshalKey(demoConf.ConfName(), &demoConf)
}

// 初始化完成
func (*stCompose) DemoDone() error {
	return nil
}

// 资源回收
func (*stCompose) DemoRecover() error {
	return nil
}
