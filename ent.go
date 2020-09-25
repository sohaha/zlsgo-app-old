// +build ignore

package main

import (
	"strings"

	e "github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zlog"
)

func main() {
	// global.Read(false)
	for _, v := range []string{
		zfile.RealPath("./ent/schema"),
		zfile.RealPath("./schema"),
	} {
		if !zfile.DirExist(v) {
			continue
		}
		packageName := "model"
		err := e.Generate(v, &gen.Config{
			Header:  "// 🙅🏻🙅🏻🙅🏻 自动生成的代码，尽量不要修改",
			Package: "app/" + packageName,
			Target:  zfile.RealPathMkdir("./" + packageName),
			IDType:  &field.TypeInfo{Type: field.TypeInt},
		})
		if err != nil && !strings.Contains(err.Error(), "no schema found") {
			zlog.Fatal("running ent codegen:", err)
		}
	}
}
