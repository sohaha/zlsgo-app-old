package global

import (
	"github.com/zlsgo/conf"
)

type (
	stManageConf struct {
		Remote       string
		Md5          string
		MaintainMode bool   `mapstructure:"maintain_mode"` // 维护模式
		IPWhitelist  string `mapstructure:"ip_whitelist"`  // 维护模式下，白名单
	}
)

func (*stManageConf) ConfName() string {
	return "manage"
}

var manageConf stManageConf

func (*stCompose) ManageReadConf(cfg *conf.Confhub) error {
	return cfg.Core.UnmarshalKey(demoConf.ConfName(), &demoConf)
}

func ManageConf() stManageConf {
	confLock.RLock()
	defer confLock.RUnlock()
	return manageConf
}

func (*stCompose) ManageDone() error {
	if manageConf.Remote == "" {
		manageConf.Remote = "https://resources.73zls.com/vue-admin-template.tar.gz"
	}

	return nil
}
