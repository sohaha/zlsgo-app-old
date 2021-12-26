package global

import (
	"database/sql"
	"fmt"
	"time"

	"app/conf"
	"github.com/sohaha/zlsgo/zutil"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB            *gorm.DB
	gormDriverMap = map[string]func(sqlDB *sql.DB) gorm.Dialector{}
)

func gormInit(sqlDB *sql.DB) (err error) {
	LogLevel := logger.Silent
	baseConf := conf.Base()
	databaseConf := conf.DB()
	dbType := databaseConf.Type
	if baseConf.Debug && databaseConf.Debug {
		LogLevel = zutil.IfVal(databaseConf.Debug, logger.Info, logger.Warn).(logger.LogLevel)
	}
	gormConfig := &gorm.Config{
		PrepareStmt: true,
		Logger: logger.New(
			Log,
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      LogLevel,
				Colorful:      true,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   databaseConf.Prefix,
			SingularTable: true,
		},
	}

	dialector, ok := gormDriverMap[dbType]
	if !ok {
		err = fmt.Errorf("not supported: %s", dbType)
		return
	}

	DB, err = gorm.Open(dialector(sqlDB), gormConfig)
	if err != nil {
		err = fmt.Errorf("failed to connect database, got error %w\n", err)
		return
	}

	// model.BindDB(DB)

	// tables := model.AutoMigrateTable()
	// err = DB.Migrator().AutoMigrate(tables...)
	// if err != nil {
	// 	return
	// }
	//
	// err = dbMigrateData(model.AutoMigrateData())
	return
}
