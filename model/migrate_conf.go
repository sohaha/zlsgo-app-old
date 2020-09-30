package model

import (
	"sync"

	"github.com/sohaha/zlsgo/zutil"
	"gorm.io/gorm"
)

type (
	migrate struct{}
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
		&AuthUser{},
		&AuthUserToken{},
		&AuthUserLogs{},
		&AuthUserGroup{},
		&AuthUserRules{},
		&AuthUserRulesRela{},
	}
}

var migrateData []func() (key string, exec func(db *gorm.DB) error)

func AutoMigrateData() []func() (key string, exec func(db *gorm.DB) error) {

	// 需要自动创建初始化执行的操作，key 是唯一
	// 🙅 不要修改历史数据！不要修改历史数据！不要修改历史数据！

	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "FirstAutoMigrateData", func(db *gorm.DB) error {
			db.Create(&MigrateLogs{
				Name: "AutoMigrateData",
			})
			return nil
		}
	})

	_ = zutil.RunAllMethod(&migrate{})
	return migrateData
}
