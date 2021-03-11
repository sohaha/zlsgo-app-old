package model

import (
	"sync"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
	"gorm.io/gorm"
)

type (
	migrate struct{}
)

var (
	db       *gorm.DB
	bindOnce sync.Once
	log      *zlog.Logger
)

func BindDB(currentDB *gorm.DB, l *zlog.Logger) {
	bindOnce.Do(func() {
		db = currentDB
		log = l
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
		&AuthGroupMenu{},
		&Menu{},
		&Setting{},
	}
}

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
