package ent

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	entsql "github.com/facebook/ent/dialect/sql"

	ent "app/model"
)

var entClient *ent.Client

func InitClient(db *sql.DB, dbType string) (err error) {
	drv := entsql.OpenDB(dbType, db)
	entClient = ent.NewClient(ent.Driver(drv))
	// defer Ent.Close()

	if err = entClient.Schema.Create(context.Background()); err != nil {
		err = fmt.Errorf("failed creating schema resources: %v", err)
	}
	return
}

func Client() (*ent.Client, error) {
	if entClient == nil {
		return nil, errors.New("please initialize the ent client first")
	}
	return entClient, nil
}

// func (*global.stCompose) EntcDone() {
// 	db, dbType := global.ZDB()
// 	// var err error
// 	// conf := DatabaseConf()
// 	// switch conf.DBType {
// 	// case "sqlite":
// 	// 	Ent, err = ent.Open("sqlite3", conf.Sqlite3.DSN())
// 	// case "mysql":
// 	// 	Ent, err = ent.Open("mysql", conf.MySQL.DSN())
// 	// }
// 	drv := entsql.OpenDB(dbType, db)
// 	Ent = ent.NewClient(ent.Driver(drv))
//
// 	// if err != nil {
// 	// 	Log.Fatal("failed opening connection to database:", err)
// 	// }
// 	// defer Ent.Close()
// 	if err := Ent.Schema.Create(context.Background()); err != nil {
// 		global.Log.Fatal("failed creating schema resources:", err)
// 	}
// }
