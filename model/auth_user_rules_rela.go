package model

import (
	"gorm.io/gorm"
)

type AuthUserRulesRela struct {
	RuleID    uint           `gorm:"column:rule_id;type:int(2)"`
	GroupID   uint           `json:"group_id"`
	Sort      uint16         `json:"sort"`
	Status    uint8          `gorm:"type:int(2);default:1"`
	ID        uint           `gorm:"primarykey" json:"id,omitempty"`
	CreatedAt JSONTime       `gorm:"column:create_time;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
