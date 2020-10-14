package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zstring"
)

// AuthUserToken 管理员权限密钥
type AuthUserToken struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Userid    uint           `gorm:"column:userid;type:int(11);not null;default:0;comment:管理员Id" json:"userid"`
	Token     string         `gorm:"column:token;type:varchar(255);not null;default:'';comment:token" json:"token"`
	IP        string         `gorm:"column:ip;type:varchar(20);not null;default:'';comment:登录IP" json:"ip"`
	UA        string         `gorm:"column:ua;type:text(0);comment:User Agent" json:"ua"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:1;comment:状态:1正常,2禁止" json:"status"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

// tokenKen token 加密 Key
var tokenKen = "zls"

func (t *AuthUserToken) CreateToken() (token string) {
	now := time.Now().UnixNano()
	token = fmt.Sprintf("%d|%s|%d|%s", t.Userid, t.IP, now, zstring.Rand(4))
	ecrypt, err := zstring.AesEnCryptString(token, tokenKen)
	if err != nil {
		return
	}
	t.Token = ecrypt
	tx := db.Model(&t).Create(t)
	if tx.Error != nil {
		return ""
	}
	return
}

// 获取上次登录信息
func (t *AuthUserToken) Last() (has bool) {
	db.Limit(1).Order("id DESC").Select("ip,ua,create_time").Find(&t)
	return t.ID != 0
}

func (t *AuthUserToken) UpdateStatus() {
	t.Status = 0
	db.Where(&t).Select("status").Updates(t)
}
