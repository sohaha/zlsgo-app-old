package model

import (
	"time"

	"gorm.io/gorm"
)

// MigrateLogs migrate log
type MigrateLogs struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

func (m MigrateLogs) Exist() (exist bool) {
	res := db.Where(m).Select("id").First(&MigrateLogs{})
	if res.Error != nil {
		// return res.Error == gorm.ErrRecordNotFound

		return false
	}

	return res.RowsAffected == 1
}

func (m *MigrateLogs) Insert() *gorm.DB {
	return db.Create(m)
}
