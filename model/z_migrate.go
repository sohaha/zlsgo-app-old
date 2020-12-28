package model

import (
	"context"
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

// FindOne 查询单条数据
func FindOne(ctx context.Context, db *gorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Exist Check 检查数据是否存在
func Exist(ctx context.Context, db *gorm.DB) (bool, error) {
	var count int64
	result := db.WithContext(ctx).Count(&count)
	if err := result.Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
