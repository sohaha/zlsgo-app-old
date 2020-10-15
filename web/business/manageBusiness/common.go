package manageBusiness

import (
	"math"
	"sync"
)

var Mutex sync.Mutex

// 字节转kb
func Byte2Kb(byteNum int64) int64 {
	k := float64(byteNum) / 1024.00
	if k == 0 {
		return 1
	}

	return int64(math.Ceil(k))
}
