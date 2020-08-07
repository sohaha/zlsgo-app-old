// +build sqlite

package module

import (
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zutil"
	sqliteDriver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	databaseDefaultInitConf["sqlite.path"] = "./db.sqlite"
}

func (*stGormDriver) GetSqlite(conf func() *gorm.Config) {
	var err error
	gormDriverMap["sqlite"], err = gorm.Open(sqliteDriver.Open(zfile.RealPath(DatabaseConf().Sqlite3.Path)), conf())
	zutil.CheckErr(err)
}
