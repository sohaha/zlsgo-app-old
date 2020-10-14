package model

import (
	"gorm.io/gorm"
)

type AuthUserRules struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Title     string         `gorm:"column:title;type:varchar(255);not null;default:'';comment:规则名称;" json:"title"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:1;comment:状态：1正常，2禁止；标识码不支持禁止" json:"status"`
	Type      uint8          `gorm:"column:type;type:tinyint(4);not null;default:1;comment:类型：1路由，2标识码" json:"type"`
	Mark      string         `gorm:"column:mark;type:text(0);not null;comment:路由类型支持多个，使用 \n 分隔" json:"mark"`
	Methods   string         `gorm:"column:methods;type:text(0);not null;comment:请求方式：大写，多个用|分隔" json:"methods"`
	Remark    string         `gorm:"column:remark;type:varchar(255);not null;default:'';comment:备注;" json:"remark"`
	Condition string         `gorm:"column:condition;type:varchar(255);not null;default:'';comment:附加条件;" json:"condition"`
	Sort      uint16         `gorm:"column:sort;type:int(11);not null;default:0;comment:排序" json:"sort"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
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
					Title: "用户管理",
					Type:  1,
					Mark:  "/ZlsManage/UserManageApi*",
				}, {
					Title:  "系统管理权限",
					Type:   2,
					Remark: "系统管理权限",
					Mark:   "systems",
				}, {
					Title:  "用户权限",
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