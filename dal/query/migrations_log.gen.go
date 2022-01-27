// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"app/dal/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newMigrationsLog(db *gorm.DB) migrationsLog {
	_migrationsLog := migrationsLog{}

	_migrationsLog.migrationsLogDo.UseDB(db)
	_migrationsLog.migrationsLogDo.UseModel(&model.MigrationsLog{})

	tableName := _migrationsLog.migrationsLogDo.TableName()
	_migrationsLog.ALL = field.NewField(tableName, "*")
	_migrationsLog.Version = field.NewInt64(tableName, "version")
	_migrationsLog.MigrationName = field.NewString(tableName, "migration_name")
	_migrationsLog.StartTime = field.NewTime(tableName, "start_time")
	_migrationsLog.EndTime = field.NewTime(tableName, "end_time")
	_migrationsLog.Breakpoint = field.NewBool(tableName, "breakpoint")

	_migrationsLog.fillFieldMap()

	return _migrationsLog
}

type migrationsLog struct {
	migrationsLogDo

	ALL           field.Field
	Version       field.Int64
	MigrationName field.String
	StartTime     field.Time
	EndTime       field.Time
	Breakpoint    field.Bool

	fieldMap map[string]field.Expr
}

func (m migrationsLog) As(alias string) *migrationsLog {
	m.migrationsLogDo.DO = *(m.migrationsLogDo.As(alias).(*gen.DO))

	m.ALL = field.NewField(alias, "*")
	m.Version = field.NewInt64(alias, "version")
	m.MigrationName = field.NewString(alias, "migration_name")
	m.StartTime = field.NewTime(alias, "start_time")
	m.EndTime = field.NewTime(alias, "end_time")
	m.Breakpoint = field.NewBool(alias, "breakpoint")

	m.fillFieldMap()

	return &m
}

func (m *migrationsLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *migrationsLog) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 5)
	m.fieldMap["version"] = m.Version
	m.fieldMap["migration_name"] = m.MigrationName
	m.fieldMap["start_time"] = m.StartTime
	m.fieldMap["end_time"] = m.EndTime
	m.fieldMap["breakpoint"] = m.Breakpoint
}

func (m migrationsLog) clone(db *gorm.DB) migrationsLog {
	m.migrationsLogDo.ReplaceDB(db)
	return m
}

type migrationsLogDo struct{ gen.DO }

func (m migrationsLogDo) Debug() *migrationsLogDo {
	return m.withDO(m.DO.Debug())
}

func (m migrationsLogDo) WithContext(ctx context.Context) *migrationsLogDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m migrationsLogDo) Clauses(conds ...clause.Expression) *migrationsLogDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m migrationsLogDo) Not(conds ...gen.Condition) *migrationsLogDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m migrationsLogDo) Or(conds ...gen.Condition) *migrationsLogDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m migrationsLogDo) Select(conds ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m migrationsLogDo) Where(conds ...gen.Condition) *migrationsLogDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m migrationsLogDo) Order(conds ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m migrationsLogDo) Distinct(cols ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m migrationsLogDo) Omit(cols ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m migrationsLogDo) Join(table schema.Tabler, on ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m migrationsLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m migrationsLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m migrationsLogDo) Group(cols ...field.Expr) *migrationsLogDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m migrationsLogDo) Having(conds ...gen.Condition) *migrationsLogDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m migrationsLogDo) Limit(limit int) *migrationsLogDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m migrationsLogDo) Offset(offset int) *migrationsLogDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m migrationsLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *migrationsLogDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m migrationsLogDo) Unscoped() *migrationsLogDo {
	return m.withDO(m.DO.Unscoped())
}

func (m migrationsLogDo) Create(values ...*model.MigrationsLog) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m migrationsLogDo) CreateInBatches(values []*model.MigrationsLog, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m migrationsLogDo) Save(values ...*model.MigrationsLog) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m migrationsLogDo) First() (*model.MigrationsLog, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrationsLog), nil
	}
}

func (m migrationsLogDo) Take() (*model.MigrationsLog, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrationsLog), nil
	}
}

func (m migrationsLogDo) Last() (*model.MigrationsLog, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrationsLog), nil
	}
}

func (m migrationsLogDo) Find() ([]*model.MigrationsLog, error) {
	result, err := m.DO.Find()
	return result.([]*model.MigrationsLog), err
}

func (m migrationsLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MigrationsLog, err error) {
	buf := make([]*model.MigrationsLog, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m migrationsLogDo) FindInBatches(result *[]*model.MigrationsLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m migrationsLogDo) Attrs(attrs ...field.AssignExpr) *migrationsLogDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m migrationsLogDo) Assign(attrs ...field.AssignExpr) *migrationsLogDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m migrationsLogDo) Joins(field field.RelationField) *migrationsLogDo {
	return m.withDO(m.DO.Joins(field))
}

func (m migrationsLogDo) Preload(field field.RelationField) *migrationsLogDo {
	return m.withDO(m.DO.Preload(field))
}

func (m migrationsLogDo) FirstOrInit() (*model.MigrationsLog, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrationsLog), nil
	}
}

func (m migrationsLogDo) FirstOrCreate() (*model.MigrationsLog, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MigrationsLog), nil
	}
}

func (m migrationsLogDo) FindByPage(offset int, limit int) (result []*model.MigrationsLog, count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = m.Offset(offset).Limit(limit).Find()
	return
}

func (m migrationsLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m *migrationsLogDo) withDO(do gen.Dao) *migrationsLogDo {
	m.DO = *do.(*gen.DO)
	return m
}