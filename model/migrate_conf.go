package model

import (
	"sync"

	"github.com/sohaha/zlsgo/zlog"
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
