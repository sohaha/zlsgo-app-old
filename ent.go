// +build ignore

package main

import (
	e "github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zlog"
)

func main() {
	for _, v := range []string{
		zfile.RealPath("./ent/schema"),
		zfile.RealPath("./schema"),
	} {
		if !zfile.DirExist(v) {
			continue
		}
		err := e.Generate(v, &gen.Config{
			Header: "// 自动生成代码，不要修改🙅🏻",
			Target: zfile.RealPathMkdir("./ent-model"),
			IDType: &field.TypeInfo{Type: field.TypeInt},
		})
		if err != nil {
			zlog.Fatal("running ent codegen:", err)
		}
	}

}
