package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if (t == JSONTime{}) {
		formatted := fmt.Sprintf("\"%s\"", "")
		return []byte(formatted), nil
	} else {
		formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
		return []byte(formatted), nil
	}
}

// Value insert timestamp into mysql need this function
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type (
	Page struct {
		Total    uint `json:"total"`
		Count    uint `json:"count"`
		Curpage  uint `json:"curpage"`
		Pagesize uint `json:"pagesize"`
	}
)

func TableName(table string) string {
	return db.NamingStrategy.TableName(table)
}

// FindPage 查询分页数据
func FindPage(ctx context.Context, db *gorm.DB, pp *Page, itmes interface{}) (int64, error) {
	total, curpage, pageSize := int64(0), pp.Curpage, pp.Pagesize
	if pageSize == 0 {
		pageSize = 10
	}
	if curpage == 0 {
		curpage = 1
	}
	err := db.WithContext(ctx).Count(&total).Error
	pp.Total = uint(total)
	pp.Count = uint(math.Ceil(float64(total) / float64(pageSize)))
	if err != nil {
		return 0, err
	} else if total == 0 || curpage > pp.Count {
		return total, nil
	}
	if curpage > 0 && pageSize > 0 {
		db = db.Offset(int((curpage - 1) * pageSize)).Limit(int(pageSize))
	} else if pageSize > 0 {
		db = db.Limit(int(pageSize))
	}

	err = db.WithContext(ctx).Find(itmes).Error
	return total, err
}
