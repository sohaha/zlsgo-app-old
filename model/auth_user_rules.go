package model

import (
	"gorm.io/gorm"
)

type AuthUserRules struct {
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	title     string         `gorm:"type:varbinary(100)"`
	Status    uint8          `gorm:"type:int(2);default:1;comment:状态：1正常，2禁止；标识码不支持禁止"`
	Type      uint8          `gorm:"column:type;type:int(2);comment:类型：1路由，2标识码"`
	Mark      string         `gorm:"comment:路由类型支持多个，使用 \n 分隔" json:"mark"`
	Methods   string         `gorm:"comment:请求方式：大写，多个用|分隔" json:"methods"`
	Remark    string         `gorm:"type:varbinary(200)"`
	Condition string         `json:"condition"`
	Sort      uint16         `json:"sort"`
	CreatedAt JSONTime       `gorm:"column:create_time;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

const (
	// AuthUserRulesMethodsSeparator 请求方式分隔符
	AuthUserRulesMethodsSeparator = "|"
	// AuthUserRulesMarkSeparator 请求方式分隔符
	AuthUserRulesMarkSeparator = "\n"
	// AuthUserRulesStatusAdopt 规则状态：通过
	AuthUserRulesStatusAdopt = 1
	// AuthUserRulesStatusIntercept 规则状态：拦截
	AuthUserRulesStatusIntercept = 2
)

func (*migrate) CreateAuthUserRules() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthUserRules", func(db *gorm.DB) error {
			data := []AuthUserRules{
				{
					title: "用户管理",
					Type:  1,
					Mark:  "ZlsManage/UserManageApi*",
				}, {
					title:  "系统管理权限",
					Type:   2,
					Remark: "系统管理权限",
					Mark:   "systems",
				}, {
					title:  "用户权限",
					Type:   1,
					Remark: "用户基础权限",
					Mark:   "ZlsManage/UserApi*",
				},
			}
			db.Create(data)
			return nil
		}
	})
}
