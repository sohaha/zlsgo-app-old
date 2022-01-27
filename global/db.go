package global

import (
	"database/sql"
	"fmt"
	"strings"

	"app/conf"
	"github.com/sohaha/zdb"
	dbmysql "github.com/sohaha/zdb/Driver/mysql"
	dbpostgres "github.com/sohaha/zdb/Driver/postgres"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/sohaha/zlsgo/zvalid"
	gormmysql "gorm.io/driver/mysql"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbDriver struct{}

var (
	dbDriverMap = make(map[string]zdb.IfeConfig, 0)
	ZDB         *zdb.Engine
)

func init() {
	initMaps = append(initMaps, func() error {
		dbConf := conf.DB()
		dbType := dbConf.Type
		if dbType == "none" {
			Log.Debug(zlog.ColorTextWrap(zlog.ColorYellow, "No database"))
			return nil
		}

		if err := validDb(); err != nil {
			return err
		}

		if err := zutil.RunAllMethod(&dbDriver{}, dbConf); err != nil {
			return err
		}
		dbDriverConf, ok := dbDriverMap[dbType]
		if !ok {
			Log.Fatal("not supported:", dbType)
		}

		var err error
		ZDB, err = zdb.New(dbDriverConf)
		if err != nil {
			return fmt.Errorf("failed opening connection to database: %v", err)
		}
		// DBDsn = dbDriverConf.GetDsn()
		return gormInit(dbDriverConf.DB())
	})
}

func validDb() error {
	dbTypes := []string{"mysql", "sqlite", "postgres"}
	return zvalid.Text(conf.DB().Type, "数据库类型").Required().
		EnumString(append(dbTypes, "none"), "数据库类型暂只支持: "+strings.Join(dbTypes, ", ")).Error()
}

func (d dbDriver) Mysql(dbConf conf.DBConf) {
	dbDriverMap["mysql"] = &dbmysql.Config{Dsn: dbConf.MySQL.DSN()}
	gormDriverMap["mysql"] = func(sqlDB *sql.DB) gorm.Dialector {
		return gormmysql.New(gormmysql.Config{
			Conn: sqlDB,
		})
	}
}

func (d dbDriver) Postgres(dbConf conf.DBConf) {
	dbDriverMap["postgres"] = &dbpostgres.Config{Dsn: dbConf.Postgres.DSN()}
	gormDriverMap["postgres"] = func(sqlDB *sql.DB) gorm.Dialector {
		return gormpostgres.New(gormpostgres.Config{
			Conn: sqlDB,
		})
	}
}
