package manageBusiness

import (
	"app/model"
	"github.com/sohaha/zlsgo/zcache"
	"github.com/sohaha/zlsgo/ztype"
)

const SETTING_CACHE_KEY = "cache_setting"
const SETTING_VALUES = "GetSetting"

var SettingFields = [4]string{"sitename", "domain", "login_expire_time", "login_mode"}

func SettingValues() interface{} {
	settingCache := zcache.New(SETTING_CACHE_KEY)
	reRes, err := settingCache.Get(SETTING_VALUES)
	if err != nil {
		re := map[string]interface{}{}
		for _, val := range (&model.Setting{}).SettingValues() {
			re[val.Varname] = val.Value
		}
		settingCache.Set(SETTING_VALUES, re, 3600)
	}

	reRes, _ = settingCache.Get(SETTING_VALUES)

	return reRes
}

// 0(false):单端登录 其余(true):多地登录
func IsMultipleLogins() bool {
	key := "login_mode"
	for k, v := range SettingValues().(map[string]interface{}) {
		if key == k {
			return v.(string) == "1"
		}
	}

	return false
}

func LoginExpireTime() int {
	key := "login_expire_time"
	for k, v := range SettingValues().(map[string]interface{}) {
		if key == k {
			return ztype.ToInt(v)
		}
	}

	return 3600
}
