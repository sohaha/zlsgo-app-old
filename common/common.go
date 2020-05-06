package common

import (
	"app/conf"
	"github.com/sohaha/zlsgo/zcache"
)

//noinspection GoVetCopyLock,GoUnusedGlobalVariable
var (
	Log   = conf.Log
	Cache = zcache.New("app")
)
