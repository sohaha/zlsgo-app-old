package model

import (
	"sync"

	"xorm.io/builder"
	"xorm.io/core"
	"xorm.io/xorm"
	xormlog "xorm.io/xorm/log"

	"app/conf"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/zutil"
	"github.com/sohaha/zlsgo/zvalid"

	_ "github.com/go-sql-driver/mysql"
)

type dbLogger struct {
	showSQL bool
	logger  *zlog.Logger
}

var (
	initDb sync.Once
	Db     *xorm.Engine
	_      xormlog.ContextLogger = &dbLogger{}
)

func Init() {
	initDb.Do(func() {
		var dsn string
		var dbType = conf.Db().DBType
		zutil.CheckErr(validDb(), true)
		switch dbType {
		case "mysql":
			dsn = conf.Db().Drive.MySQL.DSN()
		case "sqlite3":
			dsn = conf.Db().Drive.Sqlite3.DSN()
		case "none":
			return
		default:
			conf.Log.Fatal("unknown db type:", dbType)
		}
		engine, err := xorm.NewEngine(dbType, dsn)
		zutil.CheckErr(err, true)
		if conf.Db().Prefix != "" {
			tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, conf.Db().Prefix)
			engine.SetTableMapper(tbMapper)
		}
		engine.TZLocation = ztime.GetTimeZone()
		Db = engine
		engine.SetLogger(&dbLogger{
			logger: conf.Log,
		})
		if conf.Db().Debug {
			Db.ShowSQL(true)
		}
		if !conf.Db().DisableAutoMigrate {
			dbAutoMigrate()
		}
	})
}

func dbAutoMigrate() {
	m := &migrate{}
	_ = zutil.RunAllMethod(m)
	err := Db.Sync2(m.tables...)
	zutil.CheckErr(err, true)
}

func validDb() error {
	return zvalid.Text(conf.Db().DBType, "数据库类型").Required().
		EnumString([]string{
			"none", "mysql", "sqlite3",
		}, "数据库类型暂只支持: mysql, sqlite3").Error()

}

func (l *dbLogger) Warnf(format string, v ...interface{}) {
	l.logger.Debugf(format+"\n", v...)
}

func (l *dbLogger) BeforeSQL(_ xormlog.LogContext) {}

func (l *dbLogger) AfterSQL(ctx xormlog.LogContext) {
	fullSqlStr, err := builder.ConvertToBoundSQL(ctx.SQL, ctx.Args)
	if err != nil {
		l.logger.Errorf("[SQL] %v %v - %v\n", ctx.SQL, ctx.Args, ctx.ExecuteTime)
	} else {
		l.logger.Infof("[SQL] %s - %v\n", fullSqlStr, ctx.ExecuteTime)
	}
}

func (l *dbLogger) Debugf(format string, v ...interface{}) {
	l.logger.Debugf(format+"\n", v...)
}

func (l *dbLogger) Errorf(format string, v ...interface{}) {
	l.logger.Debugf(format+"\n", v...)
}

func (l *dbLogger) Infof(format string, v ...interface{}) {
	l.logger.Debugf(format+"\n", v...)
}

func (l *dbLogger) Level() xormlog.LogLevel {
	return zutil.IfVal(conf.Base().Debug, xormlog.LOG_DEBUG, xormlog.LOG_ERR).(xormlog.LogLevel)
}

func (l *dbLogger) SetLevel(ll xormlog.LogLevel) {

}

func (l *dbLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		l.showSQL = true
		return
	}
	l.showSQL = show[0]
}

func (l *dbLogger) IsShowSQL() bool {
	return l.showSQL
}
