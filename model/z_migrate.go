package model

import (
	"context"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// MigrateLogs migrate log
type MigrateLogs struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

func (m MigrateLogs) Exist() (exist bool) {
	res := db.Where(m).Select("id").First(&MigrateLogs{})
	if res.Error != nil {
		// return res.Error == gorm.ErrRecordNotFound

		return false
	}

	return res.RowsAffected == 1
}

func (m *MigrateLogs) Insert() *gorm.DB {
	return db.Create(m)
}

// FindOne 查询单条数据
func FindOne(ctx context.Context, db *gorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
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

//
// // PaginationParam 分页查询条件
// type PaginationParam struct {
// 	Pagination bool `form:"-"` // 是否使用分页查询
// 	OnlyCount  bool `form:"-"` // 是否仅查询count
// 	Current    uint `form:"current,default=1"`
// 	PageSize   uint `form:"pageSize,default=10" binding:"max=100"`
// }
//
// // PaginationResult 分页查询结果
// type PaginationResult struct {
// 	Total    int64 `json:"total"`
// 	Current  uint  `json:"current"`
// 	PageSize uint  `json:"pageSize"`
// }
//
// // GetCurrent 获取当前页
// func (a PaginationParam) GetCurrent() uint {
// 	return a.Current
// }
//
// // GetPageSize 获取页大小
// func (a PaginationParam) GetPageSize() uint {
// 	pageSize := a.PageSize
// 	if a.PageSize == 0 {
// 		pageSize = 10
// 	}
// 	return pageSize
// }
//
// // WrapFindPage 包装带有分页的查询
// func WrapFindPage(ctx context.Context, db *gorm.DB, pp PaginationParam, out interface{}) (*PaginationResult, error) {
// 	if pp.OnlyCount {
// 		var count int64
// 		err := db.Count(&count).Error
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &PaginationResult{Total: count}, nil
// 	} else if !pp.Pagination {
// 		err := db.Find(out).Error
// 		return nil, err
// 	}
//
// 	total, err := FindPage(ctx, db, pp, out)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return &PaginationResult{
// 		Total:    total,
// 		Current:  pp.GetCurrent(),
// 		PageSize: pp.GetPageSize(),
// 	}, nil
// }
