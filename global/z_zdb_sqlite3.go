// +build sqlite

package global

import (
	dbsqlite3 "github.com/sohaha/zdb/Driver/sqlite3"
)

func (d driverMap) Sqlite3() {
	conf := DatabaseConf()
	dbDriverMap["sqlite"] = &dbsqlite3.Config{
		Dsn: conf.Sqlite3.DSN(),
	}
}
