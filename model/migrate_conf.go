package model

import (
	"sync"

	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	bindOnce sync.Once
)

func BindDB(currentDB *gorm.DB) {
	bindOnce.Do(func() {
		db = currentDB
	})
}

func AutoMigrateTable() []interface{} {
	return []interface{}{
		// 需要自动创建的表
		&MigrateLogs{},
	}
}

func AutoMigrateData() map[string]func(DB *gorm.DB) error {
	data := map[string]func(DB *gorm.DB) error{}

	// 需要自动创建初始化执行的操作，key 是唯一
	// 🙅不要修改历史数据！不要修改历史数据！不要修改历史数据！

	data["initDemo"] = func(DB *gorm.DB) error {
		return nil
	}

	return data
}
