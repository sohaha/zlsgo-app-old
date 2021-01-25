package model

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zstring"
)

const (
	TokenNormal        uint8  = 1
	TokenDisabled      uint8  = 2
	TokenEffectiveTime string = "1h" // 有效时间
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
	tx := db.Begin()

	err := tx.Model(&t).Create(t).Error
	if err != nil {
		tx.Rollback()
		return ""
	}

	now := time.Now().UnixNano()
	token = fmt.Sprintf("%d|%s|%d|%d|%s", t.Userid, t.IP, t.ID, now, zstring.Rand(4))
	ecrypt, err := zstring.AesEnCryptString(token, tokenKen)
	if err != nil {
		return ""
	}
	t.Token = ecrypt
	err = tx.Model(&t).Select("token").Updates(AuthUserToken{Token: t.Token}).Error
	if err != nil {
		return ""
	}
	tx.Commit()

	return
}

func (t *AuthUserToken) DeToken() (token string, err error) {
	token, err = zstring.AesDeCryptString(t.Token, tokenKen)
	return
}

// 获取上次登录信息
func (t *AuthUserToken) Last() (has bool) {
	db.Limit(1).Order("id DESC").Select("ip,ua,create_time").Find(&t)
	return t.ID != 0
}

func (t *AuthUserToken) UpdateStatus() {
	t.Status = 2
	db.Where(&t).Select("status").Updates(t)
}

func (t *AuthUserToken) ClearAllToken() error {
	uRes := db.Model(&AuthUserToken{}).Select("status").Where("status = ? and userid = ?", 1, t.Userid).Updates(&AuthUserToken{Status: 2})
	if uRes.RowsAffected < 1 {
		return errors.New("服务繁忙，请重试")
	}

	return nil
}

func (t *AuthUserToken) LoginModeTrue() error {
	var res []AuthUserToken
	db.Select("MAX(id) as id, userid").Where("userid != ? and status = ?", t.Userid, TokenNormal).Group("userid").Find(&res)
	idMap := []uint{t.ID}
	for _, v := range res {
		idMap = append(idMap, v.ID)
	}
	if uRes := db.Model(AuthUserToken{}).Where("status = ? and id NOT IN ?", TokenNormal, idMap).Updates(AuthUserToken{Status: TokenDisabled}); uRes.Error != nil {
		return errors.New("更新失败")
	}

	return nil
}

func (t *AuthUserToken) TokenRules() (string, error) {
	deToken, err := t.DeToken()
	if err != nil {
		return "", err
	}
	if strings.Count(deToken, "|") != 4 {
		return "", errors.New("token错误")
	}

	return deToken, nil
}

func (t *AuthUserToken) SelectUser() *AuthUserToken {
	db.Where("id = ?", t.ID).First(t)

	return t
}

func (t *AuthUserToken) UpdateUserToken() {
	db.Model(&AuthUserToken{}).Select("update_time, status").Where("userid = ? and id != ? and status != ?", t.Userid, t.ID, TokenDisabled).Updates(AuthUserToken{Status: TokenDisabled})
}
