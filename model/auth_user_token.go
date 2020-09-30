package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zstring"
)

// AuthUserToken 管理员权限密钥
type AuthUserToken struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	Userid    uint           `json:"-"`
	Token     string         `json:"-"`
	IP        string         `json:"ip"`
	UA        string         `json:"ua"`
	Status    uint8          `gorm:"default:1" json:"status,omitempty"`
	CreatedAt JSONTime       `gorm:"column:create_time;" json:"create_time,omitempty"`
	UpdatedAt JSONTime       `gorm:"column:update_time;" json:"update_time,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
