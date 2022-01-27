package conf

import (
	"github.com/sohaha/zlsgo/zstring"
)

type (
	ManageConf struct {
		Key    string
		Expire int
	}
)

func init() {
	allConf["manage"] = &ManageConf{
		Key:    zstring.Rand(4),
		Expire: 7200,
	}
}

// Manage Manage配置
func Manage() ManageConf {
	conf := getConf(func() interface{} {
		conf, ok := allConf["manage"]
		if !ok {
			return DemoConf{}
		}
		return *conf.(*ManageConf)
	})

	return conf.(ManageConf)
}
