package global

import (
	"database/sql"

	"github.com/sohaha/zdb"
	dbmysql "github.com/sohaha/zdb/Driver/mysql"

	"app/schema"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zutil"

	// "github.com/sohaha/zdb/Driver/postgres"

	"app/global/ent"
)

type driverMap struct{}

var (
	DB          *sql.DB
	dbType      string
	dbDriverMap = map[string]zdb.IfeConfig{}
)

func (*stCompose) DBDone() {
	err := zutil.RunAllMethod(&driverMap{})
	zutil.CheckErr(err)

	conf := DatabaseConf()
	dbType = conf.DBType

	dbConf, ok := dbDriverMap[dbType]
	if !ok {
		Log.Fatal("not supported:", dbType)
	}

	_, err = zdb.New(dbConf)
	if err != nil {
		Log.Fatal("failed opening connection to database:", err)
	}

	DB = dbConf.DB()

	schema.Prefix = conf.Prefix
	zlog.Debug(schema.Prefix)
	err = ent.InitClient(DB, dbType)
	zutil.CheckErr(err)
}

func (d driverMap) Mysql() {
	conf := DatabaseConf()
	dbDriverMap["mysql"] = &dbmysql.Config{Dsn: conf.MySQL.DSN()}
}
