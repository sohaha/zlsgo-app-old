// +build ignore

package main

import (
	e "github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zlog"
	"log"
)

func main() {
	log.Println(66)
	for _, v := range []string{
		zfile.RealPath("./ent/schema"),
	} {
		if !zfile.DirExist(v) {
			continue
		}
		err := e.Generate(v, &gen.Config{
			Header: "// Your Custom Header",
			IDType: &field.TypeInfo{Type: field.TypeInt},
		})
		if err != nil {
			zlog.Fatal("running ent codegen:", err)
		}
	}

}
