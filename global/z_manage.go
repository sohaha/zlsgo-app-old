package global

import (
	"github.com/sohaha/gconf"
)

type (
	stManageConf struct {
		Path   string
		Remote string
		Md5    string
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
	if manageConf.Remote == "" {
		manageConf.Remote = "https://resources.73zls.com/vue-admin-template.tar.gz"
	}
	if manageConf.Path == "" {
		manageConf.Path = "resource/manage/"
	}

	return nil
}
