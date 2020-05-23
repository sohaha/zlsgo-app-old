package model

import (
	"bytes"
	"time"

	"github.com/sohaha/zlsgo/zvalid"
)

type (
	jsonTime time.Time
	migrate  struct {
		tables []interface{}
	}
)

func (j jsonTime) String() string {
	t := time.Time(j)
	if t.IsZero() {
		return "0000-00-00 00:00:00"
	}
	return t.Format("2006-01-02 15:04:05")
}

func (j jsonTime) MarshalJSON() ([]byte, error) {
	res := bytes.NewBufferString("\"")
	res.WriteString(j.String())
	res.WriteString("\"")
	return res.Bytes(), nil
}

func getValidRule(rules *map[string]zvalid.Engine, fields ...string) map[string]zvalid.Engine {
	validRules := *rules
	if len(fields) == 0 {
		return validRules
	}
	rule := make(map[string]zvalid.Engine)
	for _, k := range fields {
		if v, ok := validRules[k]; ok {
			rule[k] = v
		}
	}
	return rule
}
