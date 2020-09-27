// +build sqlite

package global

import (
	"database/sql"

	"github.com/sohaha/zlsgo/zfile"
	gormsqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (*dbDriver) GormSqlite(conf stDatabaseConf) {
	gormDriverMap["sqlite"] = func(sqlDB *sql.DB) gorm.Dialector {
		return gormsqlite.Open(zfile.RealPath(DatabaseConf().Sqlite.Path))
	}
}
