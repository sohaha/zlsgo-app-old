package main

import (
	"strings"

	"app/conf"
	"app/global"
	"github.com/sohaha/zlsgo/zfile"
	"gorm.io/gen"
)

func main() {
	conf.Init(zfile.RealPath("../../" + conf.FileName))
	global.Init()
	dbConf := conf.DB()
	g := gen.NewGenerator(gen.Config{
		OutPath: "../../dal/query",
	})
	g.WithFileNameStrategy(func(tableName string) string {
		tablePrefix := dbConf.Prefix
		return strings.TrimPrefix(tableName, tablePrefix)
	})

	g.UseDB(global.DB)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
