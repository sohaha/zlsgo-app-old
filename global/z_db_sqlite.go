// +build sqlite

package global

import (
	dbsqlite3 "github.com/sohaha/zdb/Driver/sqlite3"
)

func init() {
	databaseDefaultInitConf["sqlite.path"] = "./db.sqlite"
}

func (d dbDriver) Sqlite3(conf stDatabaseConf) {
	dbDriverMap["sqlite"] = &dbsqlite3.Config{
		Dsn: conf.Sqlite.DSN(),
	}
}
