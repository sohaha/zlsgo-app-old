package global

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
	gormmysql "gorm.io/driver/mysql"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"app/model"
)

var (
	DB            *gorm.DB
	gormDriverMap = map[string]func(sqlDB *sql.DB) gorm.Dialector{}
)

func (*dbDriver) GormMysql(conf stDatabaseConf) {
	gormDriverMap["mysql"] = func(sqlDB *sql.DB) gorm.Dialector {
		return gormmysql.New(gormmysql.Config{
			Conn: sqlDB,
		})
	}
}

func (*dbDriver) GormPostgres(conf stDatabaseConf) {
	gormDriverMap["postgres"] = func(sqlDB *sql.DB) gorm.Dialector {
		return gormpostgres.New(gormpostgres.Config{
			Conn: sqlDB,
		})
	}
}

func gormInit(dbType string, sqlDB *sql.DB) (err error) {
	LogLevel := logger.Silent
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

	model.BindDB(DB, Log)

	tables := model.AutoMigrateTable()
	err = DB.Migrator().AutoMigrate(tables...)
	if err != nil {
		return
	}

	err = dbMigrateData(model.AutoMigrateData())
	return
}

func dbMigrateData(data []func() (key string, exec func(db *gorm.DB) error)) (err error) {
	var ignore []string

	for _, d := range data {
		k, v := d()

		currentMigrate := &model.MigrateLogs{Name: k}
		if currentMigrate.Exist() {
			ignore = append(ignore, k)
			continue
		}
		err = v(DB)
		if err != nil {
			err = fmt.Errorf("migrate Error: [ %s ] %s", k, err.Error())
			return
		}
		err = currentMigrate.Insert().Error
	}

	if len(ignore) > 0 {
		Log.Warnf(zlog.ColorTextWrap(zlog.ColorWhite, "ignore migrate: [ %s ]\n"), strings.Join(ignore, ","))
	}
	return
}
