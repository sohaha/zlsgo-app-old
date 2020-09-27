package model

import (
	"gorm.io/gorm"
)

type AuthUserRulesRela struct {
	gorm.Model
	RuleID  uint `gorm:"column:rule_id;type:int(2)"`
	GroupID uint
	sort    uint16
	Status  uint8 `gorm:"type:int(2);default:1"`
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
				},
				{
					RuleID:  3,
					GroupID: 0,
					Status:  1,
				},
			}
			db.Create(data)
			return nil
		}
	})
}
