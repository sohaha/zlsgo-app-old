//go:build sqlite
// +build sqlite

package global

import (
	"database/sql"

	"app/conf"
	dbsqlite3 "github.com/sohaha/zdb/Driver/sqlite3"
	"github.com/sohaha/zlsgo/zfile"
	gormsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (d dbDriver) Sqlite3(dbConf conf.DBConf) {
	dbDriverMap["sqlite"] = &dbsqlite3.Config{
		Dsn: dbConf.Sqlite.DSN(),
	}
	gormDriverMap["sqlite"] = func(sqlDB *sql.DB) gorm.Dialector {
		return gormsqlite.Open(zfile.RealPath(dbConf.Sqlite.Path))
	}
}
