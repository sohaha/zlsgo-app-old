package model

import (
	"gorm.io/gorm"
)

// AuthUserRules 角色权限规则对应
type AuthUserRulesRela struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	RuleID    uint           `gorm:"column:rule_id;type:int(11);not null;default:0;comment:规则id" json:"rule_id"`
	GroupID   uint           `gorm:"column:group_id;type:int(11);not null;default:0;comment:角色id" json:"group_id"`
	Sort      uint16         `gorm:"column:sort;type:int(11);not null;default:0;comment:排序" json:"sort"`
	Status    uint8          `gorm:"column:status;type:tinyint(4);not null;default:1;comment:状态:1正常，2禁止，3忽略" json:"status"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

func (*migrate) CreateAuthUserRulesRela() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthUserRulesRela", func(db *gorm.DB) error {
			data := []AuthUserRulesRela{
				{
					RuleID:  1,
					GroupID: 1,
					Status:  1,
				}, {
					RuleID:  2,
					GroupID: 1,
					Status:  1,
				}, {
					RuleID:  3,
					GroupID: 1,
					Status:  1,
				}, {
					RuleID:  3,
					GroupID: 2,
					Status:  1,
				},
			}
			db.Create(data)
			return nil
		}
	})
}

type GroupIntegration struct {
	RuleIds    []uint `json:"rule_ids"`
	BanRuleIds []uint `json:"ban_rule_ids"`
	UserCount  int64  `json:"user_count"`
}

// 整合角色相关数据
func (rr *AuthUserRulesRela) Integration() *GroupIntegration {
	lists := []AuthUserRulesRela{}
	db.Where("group_id = ?", rr.GroupID).Find(&lists)

	groupIntegration := &GroupIntegration{RuleIds: []uint{}, BanRuleIds: []uint{}, UserCount: 0}
	for _, v := range lists {
		switch true {
		case v.Status == 1:
			groupIntegration.RuleIds = append(groupIntegration.RuleIds, v.RuleID)
			break
		case v.Status == 2:
			groupIntegration.BanRuleIds = append(groupIntegration.BanRuleIds, v.RuleID)
			break
		}
	}

	db.Model(&AuthUser{}).Where("group_id = ?", rr.GroupID).Count(&groupIntegration.UserCount)

	return groupIntegration
}
