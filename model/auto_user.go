package model

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zvalid"
)

// AuthUser 管理员
type AuthUser struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	Username  string         `json:"username"`
	Password  string         `json:"-"`
	Key       string         `json:"-"`
	Nickname  string         `json:"nickname"`
	Email     string         `json:"email"`
	Avatar    string         `json:"avatar"`
	Status    uint8          `gorm:"default:1" json:"status"`
	GroupID   uint           `gorm:"column:group_id" json:"group_id"`
	IsSuper   bool           `gorm:"column:is_super;default:0;" json:"is_super"`
	CreatedAt JSONTime       `gorm:"column:create_time;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *AuthUser) Lists(pp *Page) (users []AuthUser) {
	_, _ = FindPage(context.Background(), db.Where(u).Model(u).Order("id desc"), pp, &users)
	return
}

func (*migrate) CreateAuthUser() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		// 默认管理员密码
		password := "admin666" // zstring.Rand(6)
		encryptPassword, _ := zvalid.Text(password).EncryptPassword().String()
		return "CreateAuthUser", func(db *gorm.DB) error {
			db.Create(&AuthUser{
				Username: "admin",
				Password: encryptPassword,
				Key:      "",
				Nickname: "管理员",
				Email:    "admin@qq.com",
				Avatar:   "",
				GroupID:  1,
				Status:   1,
				IsSuper:  true,
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

func (u *AuthUser) TokenToInfo(t *AuthUserToken) {
	t.Status = 1
	db.Where(t).Limit(1).Find(&t)
	if t.Userid != 0 {
		db.Where(&AuthUser{ID: t.Userid}).Limit(0).Find(&u)
	}
}

func (u *AuthUser) Login(ip string, ua string) (string, error) {
	password := u.Password
	db.Model(&u).Where(&AuthUser{Username: u.Username}).Limit(1).Find(&u)
	if u.ID == 0 {
		return "", errors.New("用户不存在")
	}
	if err := zvalid.Text(password, "用户密码").CheckPassword(u.Password).Error(); err != nil {
		return "", err
	}

	t := AuthUserToken{
		IP:     ip,
		UA:     ua,
		Userid: u.ID,
	}
	token := t.CreateToken()
	if token == "" {
		return "", errors.New("创建 Token 失败")
	}
	return t.Token, nil
}
