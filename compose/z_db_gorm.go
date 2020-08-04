package compose

import (
	"fmt"
	"strings"
	"time"

	mysqlDriver "gorm.io/driver/mysql"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"app/model"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
)

type (
	stGormDriver struct{}
)

var (
	DB            *gorm.DB
	gormDriverMap = map[string]*gorm.DB{}
)

func init() {
}

func (*stCompose) GormDone() {
	dbType := DatabaseConf().DBType
	zutil.CheckErr(validDb())

	LogLevel := logger.Silent
	if baseConf.Debug && databaseConf.Debug {
		LogLevel = zutil.IfVal(databaseConf.Debug, logger.Info, logger.Warn).(logger.LogLevel)
	}

	err := zutil.RunAllMethod(&stGormDriver{}, func() *gorm.Config {
		return &gorm.Config{
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
	})
	zutil.CheckErr(err)

	var ok bool
	DB, ok = gormDriverMap[dbType]
	if !ok {
		if dbType == "none" {
			Log.Debug(zlog.ColorTextWrap(zlog.ColorYellow, "No database"))
			return
		}
		zutil.CheckErr(fmt.Errorf("not supported: %s", dbType))
		return
	}

	sqlDB, err := DB.DB()
	if err == nil {
		err = sqlDB.Ping()
	}
	if err != nil {
		err = fmt.Errorf("failed to connect database, got error %w\n", err)
	}
	zutil.CheckErr(err)

	model.BindDB(DB)
	zutil.CheckErr(DB.AutoMigrate(model.AutoMigrateTable()...))
	zutil.CheckErr(dbMigrateData(model.AutoMigrateData()))
}

func dbMigrateData(data map[string]func(DB *gorm.DB) error) (err error) {
	var ignore []string
	// 会是随机
	for k, v := range data {
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
		currentMigrate.Insert()
	}

	if len(ignore) > 0 {
		Log.Warnf(zlog.ColorTextWrap(zlog.ColorWhite, "ignore migrate: [ %s ]\n"), strings.Join(ignore, ","))
	}
	return
}

func (*stGormDriver) GetMysql(conf func() *gorm.Config) {
	var err error
	gormDriverMap["mysql"], err = gorm.Open(mysqlDriver.Open(DatabaseConf().MySQL.DSN()), conf())
	zutil.CheckErr(err)
}

func (*stGormDriver) GetPostgres(conf func() *gorm.Config) {
	var err error
	gormDriverMap["postgres"], err = gorm.Open(postgresDriver.Open(DatabaseConf().Postgres.DSN()), conf())
	zutil.CheckErr(err)
}
