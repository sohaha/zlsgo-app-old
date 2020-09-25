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

func MustClient() (*ent.Client, error) {
	if entClient == nil {
		return nil, errors.New("please initialize the ent client first")
	}
	return entClient, nil
}

func Client() *ent.Client {
	return entClient
}
