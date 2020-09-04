package global

import (
	"app/global/entc"
)

type (
	stEntcConf struct {
		Host     string
		Port     int
		Password string
		DBNumber int `mapstructure:"db"`
	}
)

func (*stCompose) EntcDone() {
	Log.Warn("entc")
	entc.Generate()
}
