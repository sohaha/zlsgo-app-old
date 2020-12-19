package global

import (
	"github.com/sohaha/gconf"
)

type (
	stManageConf struct {
		Path string
		Gz   string
	}
)

func (*stManageConf) ConfName() string {
	return "manage"
}

var manageConf stManageConf

func (*stCompose) ManageReadConf(cfg *gconf.Confhub) error {
	return cfg.Core.UnmarshalKey(demoConf.ConfName(), &demoConf)
}

func ManageConf() stManageConf {
	confLock.RLock()
	defer confLock.RUnlock()
	return manageConf
}

func (*stCompose) ManageDone() error {
	if manageConf.Gz == "" {
		manageConf.Gz = "http://127.0.0.1:1234/z.tar.gz"
	}
	if manageConf.Path == "" {
		manageConf.Path = "resource/manage/"
	}

	return nil
}
