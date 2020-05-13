package common

import (
	"github.com/sohaha/wechat"
)

// Wx 微信
//noinspection GoUnusedGlobalVariable
var (
	Wx     *wechat.Engine
	WxOpen *wechat.Engine
	WxQy   *wechat.Engine
)

func init() {
	wechat.Debug()
	_ = wechat.LoadCacheData("wechat.json")
	Wx = wechat.New(&wechat.Mp{
		AppID:     "wx6a24b584b45b6791",
		AppSecret: "cc31573bfa7af4cdc2ba327357af9234",
		Token:     "wx",
	})
	WxOpen = wechat.New(&wechat.Open{
		AppID:          "wx1a01b1c7bbf39ac1",
		AppSecret:      "a9de468f1b8b9268c921a0814da09187",
		EncodingAesKey: "dt666dt666dt666dt666dt666dt666dt666dt666dt6",
	})
	WxQy = wechat.New(&wechat.Qy{
		CorpID:         "wwc2ebdec1444eab71",
		Secret:         "_IoAuf1cxNhcnvn6cLPhMMeyCAA6kPJ-jZHkWadE4_U",
		Token:          "wxwx",
		EncodingAesKey: "wHG3DlQBPSsDPs7zYM2zahtqY6Y35XKoQECfJtYx1UL",
	})
}

func SaveWxCacheData() (string, error) {
	return wechat.SaveCacheData("wechat.json")
}
