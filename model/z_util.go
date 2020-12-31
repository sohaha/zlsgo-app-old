package model

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	// JSONTime format json time field by myself
	JSONTime struct {
		time.Time
	}
	Page struct {
		Total    uint `json:"total"`
		Count    uint `json:"count"`
		Curpage  uint `json:"curpage"`
		Pagesize uint `json:"pagesize"`
	}
)

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

// WithTableName 表名包裹表前缀
func WithTableName(table string) string {
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

// FindOne 查询单条数据
func FindOne(ctx context.Context, db *gorm.DB, out interface{}) (bool, error) {
	result := db.WithContext(ctx).First(out)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Check 检查数据是否存在
func Exist(ctx context.Context, db *gorm.DB) (bool, error) {
	var count int64
	result := db.WithContext(ctx).Count(&count)
	if err := result.Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// OrderFieldFunc 排序字段转换函数
type OrderFieldFunc func(string) string

// OrderDirection 排序方向
type OrderDirection int

const (
	// OrderByASC 升序排序
	OrderByASC OrderDirection = 1
	// OrderByDESC 降序排序
	OrderByDESC OrderDirection = 2
)

// OrderField 排序字段
type OrderField struct {
	Key       string         // 字段名
	Direction OrderDirection // 排序方向
}

// ParseOrder 解析排序字段
func ParseOrder(items []*OrderField, handle ...OrderFieldFunc) string {
	orders := make([]string, len(items))

	for i, item := range items {
		key := item.Key
		if len(handle) > 0 {
			key = handle[0](key)
		}

		direction := "ASC"
		if item.Direction == OrderByDESC {
			direction = "DESC"
		}
		orders[i] = fmt.Sprintf("%s %s", key, direction)
	}

	return strings.Join(orders, ",")
}
