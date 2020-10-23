package model

import (
	"gorm.io/gorm"
)

// AuthGroupMenu 角色菜单对应表
type AuthGroupMenu struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	GroupID   uint8          `gorm:"column:groupid;type:tinyint(4);not null;default:0;comment:角色Id;" json:"groupid"`
	Menu      string         `gorm:"column:menu;type:varchar(255);not null;default:'';comment:菜单;" json:"menu"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

func (*migrate) CreateAuthGroupMenu() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthGroupMenu", func(db *gorm.DB) error {
			db.Create([]AuthGroupMenu{
				{
					GroupID: 1,
					Menu:    "1,2,3,4,5,6,7,8,9,10,11",
				},
				{
					GroupID: 2,
					Menu:    "1",
				},
			})
			return nil
		}
	})
}
