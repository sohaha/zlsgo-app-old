package global

import (
	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zfile"

	"github.com/sohaha/wechat"
)

type (
	stWxConf struct {
		Debug bool
		Mp    wechat.Mp
		Open  wechat.Open
		Qy    wechat.Qy
	}
)

const wxFile = "wechat.json"

// 根据项目实际情况使用对应的
//noinspection GoUnusedGlobalVariable
var (
	WxMp   *wechat.Engine
	WxOpen *wechat.Engine
	WxQy   *wechat.Engine
	wxConf stWxConf
)

func (*stWxConf) ConfName(key ...string) string {
	if len(key) > 0 {
		return "wx." + key[0]
	}
	return "wx"
}

func (*stCompose) WxDefaultConf(cfg *gconf.Confhub) {
	conf := map[string]interface{}{}
	conf["debug"] = true

	// 如果不需要可以删除
	conf["mp"] = map[string]interface{}{
		"appid":     "",
		"appsecret": "",
		"token":     "",
	}

	conf["open"] = map[string]interface{}{

	}

	conf["qy"] = map[string]interface{}{

	}

	cfg.SetDefault(wxConf.ConfName(), conf)
}

func (*stCompose) WxReadConf(cfg *gconf.Confhub) error {
	return cfg.Core.UnmarshalKey(wxConf.ConfName(), &wxConf)
}

func (*stCompose) WxDone() {
	if wxConf.Debug && BaseConf().Debug {
		wechat.Debug()
	}
	_ = wechat.LoadCacheData(zfile.RealPath(wxFile))
	WxMp = wechat.New(&wxConf.Mp)
	WxOpen = wechat.New(&wxConf.Open)
	WxQy = wechat.New(&wxConf.Qy)
}

func (*stCompose) WxRecover() {
	_, _ = wechat.SaveCacheData(zfile.RealPath(wxFile))
}
