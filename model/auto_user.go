package model

import (
	"context"
	"errors"

	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/zvalid"
	"gorm.io/gorm"
)

// AuthUser 管理员
type AuthUser struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Key      string `json:"key"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Status   uint8  `gorm:"default:1" json:"status"`
	GroupID  uint   `gorm:"index" json:"group_id"`
}

func (*migrate) CreateAuthUser() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		password := zstring.Rand(6)
		encryptPassword, _ := zvalid.Text(password).EncryptPassword().String()
		return "CreateAuthUser", func(db *gorm.DB) error {
			db.Create(&AuthUser{
				Username: "admin",
				Password: encryptPassword,
				Key:      "",
				Nickname: "管理员",
				Email:    "admin@qq.com",
				Avatar:   "",
				Status:   1,
			})
			return nil
		}
	})
}

func (u *AuthUser) Insert() error {
	has, err := u.UserExist()
	if err != nil {
		return err
	}
	if has {
		return errors.New("用户已存在")
	}
	tx := db.Create(u)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (u *AuthUser) UserExist() (bool, error) {
	where := AuthUser{}
	if u.Username != "" {
		where.Username = u.Username
	}
	if u.Email != "" {
		where.Email = u.Email
	}
	return Exist(context.Background(), db.Where(where).Model(&where))
}

func (u *AuthUser) TokenToInfo(token string) {

}
