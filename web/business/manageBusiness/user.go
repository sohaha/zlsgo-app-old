package manageBusiness

import (
	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/ztype"
	"os"
	"strings"
)

type PutUpdateSt struct {
	Id        uint   `json:"id"`
	Status    uint8  `json:"status"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
	Avatar    string `json:"avatar"`
	Remark    string `json:"remark"`
	Email     string `json:"email"`
	GroupID   uint   `json:"group_id"`
	Nickname  string `json:"nickname"`
}

type PutEditPasswordSt struct {
	OldPass string `json:"oldPass"`
	Pass    string `json:"pass"`
	Pass2   string `json:"pass2"`
	UserID  uint   `json:"userid"`
}

func IsAdmin(userid uint) int {
	cfg := gconf.New("conf.yml")
	_ = cfg.Read()
	adminCfg := cfg.Get("admin")
	if adminCfg != nil {
		userSlice := strings.Split(adminCfg.(string), ",")
		for _, user := range userSlice {
			userStr := strings.TrimSpace(user)
			if userStr == ztype.ToString(userid) {
				return 1
			}
		}
	}

	pool := map[uint]int8{
		1: 1,
	}

	_, has := pool[userid]
	if has == true {
		return 1
	}

	return 0
}

// 移动临时头像文件到指定目录上
func MvAvatar(path string, filename string) (newPath string, err error) {
	completePath := zfile.RealPathMkdir(avatarPath, true)

	newAva := completePath + filename

	err = os.Rename(zfile.RealPath(path), newAva)
	if err != nil {
		return path, nil
	}

	return "/" + zfile.SafePath(newAva), nil
}
