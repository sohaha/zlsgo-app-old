package global

import (
	"github.com/sohaha/zlsgo/zutil"
)

type initFn func() error

var (
	initMaps = make([]initFn, 0)
)

func Init() {
	for _, fn := range initMaps {
		zutil.CheckErr(fn(), true)
	}
}
