package model

import (
	"gorm.io/gorm"
)

type AuthUserRules struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt JSONTime       `gorm:"column:create_time;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	title     string         `gorm:"type:varbinary(100)"`
	Status    uint8          `gorm:"type:int(2);default:1"`
	Type      uint8          `gorm:"column:type;type:int(2)"`
	Mark      string
	Remark    string `gorm:"type:varbinary(200)"`
	Condition string
	sort      uint16
}

func (*migrate) CreateAuthUserRules() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthUserRules", func(db *gorm.DB) error {
			data := []AuthUserRules{
				{
					title: "用户管理",
					Type:  1,
					Mark:  "/ZlsManage/UserManageApi*",
				}, {
					title:  "系统管理权限",
					Type:   2,
					Remark: "系统管理权限",
					Mark:   "systems",
				}, {
					title:  "登录权限",
					Type:   1,
					Remark: "系统管理权限",
					Mark:   "/ZlsManage/UserApi/GetToken.go",
				}, {
					title:  "用户权限",
					Type:   1,
					Remark: "用户基础权限",
					Mark:   "/ZlsManage/UserApi*",
				},
			}
			db.Create(data)
			return nil
		}
	})
}
