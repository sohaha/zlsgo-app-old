package model

import (
	"github.com/sohaha/zlsgo/zutil"
	"gorm.io/gorm"
)

var migrateData []func() (key string, exec func(db *gorm.DB) error)

func AutoMigrateData() []func() (key string, exec func(db *gorm.DB) error) {

	// 需要自动创建初始化执行的操作，key 是唯一
	// 🙅 不要修改历史数据！不要修改历史数据！不要修改历史数据！

	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "FirstAutoMigrateData", // Key
			func(db *gorm.DB) error { // 数据处理
				db.Create(&MigrateLogs{
					Name: "AutoMigrateData",
				})
				return nil
			}
	})

	_ = zutil.RunAllMethod(&migrate{})

	return migrateData
}

