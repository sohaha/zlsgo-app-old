package model

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/sohaha/zlsgo/zvalid"
)

// AuthUser 管理员
type AuthUser struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Username  string         `gorm:"column:username;type:varchar(255);not null;default:'';comment:用户名;" json:"username"`
	Password  string         `gorm:"column:password;type:varchar(255);not null;default:'';comment:用户密码;" json:"-"`
	Key       string         `gorm:"column:key;type:varchar(255);not null;default:'';comment:密码盐;" json:"-"`
	Nickname  string         `gorm:"column:nickname;type:varchar(255);not null;default:'';comment:用户昵称;" json:"nickname"`
	Email     string         `gorm:"column:email;type:varchar(255);not null;default:'';comment:Email;" json:"email"`
	Remark    string         `gorm:"column:remark;type:varchar(255);not null;default:'';comment:用户简介;" json:"remark"`
	Avatar    string         `gorm:"column:avatar;type:varchar(255);not null;default:'';comment:头像;" json:"avatar"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:0;comment:状态:-1软删除,0待激活,1正常,2禁止;" json:"status"`
	GroupID   uint           `gorm:"autoIncrement;column:group_id;type:int(11);not null;default:0;comment:角色Id;" json:"group_id"`
	IsSuper   bool           `gorm:"column:is_super;default:0;" json:"is_super"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

// 默认管理员密码
const DefManagePassword = "admin666"

func (u *AuthUser) Lists(pp *Page) (users []AuthUser) {
	_, _ = FindPage(context.Background(), db.Where(u).Model(u).Order("id desc"), pp, &users)
	return
}

func (*migrate) CreateAuthUser() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		password := DefManagePassword // zstring.Rand(6)
		encryptPassword, _ := zvalid.Text(password).EncryptPassword().String()
		return "CreateAuthUser", func(db *gorm.DB) error {
			db.Create([]AuthUser{
				{
					Username: "manage",
					Password: encryptPassword,
					Key:      "",
					Nickname: "超级管理员",
					Email:    "manage@qq.com",
					Avatar:   "",
					GroupID:  1,
					Status:   1,
					IsSuper:  true,
				},
				{
					Username: "admin",
					Password: encryptPassword,
					Key:      "",
					Nickname: "管理员",
					Email:    "admin@qq.com",
					Avatar:   "",
					GroupID:  1,
					Status:   1,
					IsSuper:  false,
				},
				{
					Username: "edit",
					Password: encryptPassword,
					Key:      "",
					Nickname: "编辑",
					Email:    "edit@qq.com",
					Avatar:   "",
					GroupID:  2,
					Status:   1,
					IsSuper:  false,
				},
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