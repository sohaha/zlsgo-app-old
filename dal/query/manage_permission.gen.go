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

func newManagePermission(db *gorm.DB) managePermission {
	_managePermission := managePermission{}

	_managePermission.managePermissionDo.UseDB(db)
	_managePermission.managePermissionDo.UseModel(&model.ManagePermission{})

	tableName := _managePermission.managePermissionDo.TableName()
	_managePermission.ALL = field.NewField(tableName, "*")
	_managePermission.ID = field.NewInt32(tableName, "id")
	_managePermission.Ptype = field.NewString(tableName, "ptype")
	_managePermission.V0 = field.NewString(tableName, "v0")
	_managePermission.V1 = field.NewString(tableName, "v1")
	_managePermission.V2 = field.NewString(tableName, "v2")
	_managePermission.V3 = field.NewString(tableName, "v3")
	_managePermission.V4 = field.NewString(tableName, "v4")
	_managePermission.V5 = field.NewString(tableName, "v5")

	_managePermission.fillFieldMap()

	return _managePermission
}

type managePermission struct {
	managePermissionDo

	ALL   field.Field
	ID    field.Int32
	Ptype field.String
	V0    field.String
	V1    field.String
	V2    field.String
	V3    field.String
	V4    field.String
	V5    field.String

	fieldMap map[string]field.Expr
}

func (m managePermission) As(alias string) *managePermission {
	m.managePermissionDo.DO = *(m.managePermissionDo.As(alias).(*gen.DO))

	m.ALL = field.NewField(alias, "*")
	m.ID = field.NewInt32(alias, "id")
	m.Ptype = field.NewString(alias, "ptype")
	m.V0 = field.NewString(alias, "v0")
	m.V1 = field.NewString(alias, "v1")
	m.V2 = field.NewString(alias, "v2")
	m.V3 = field.NewString(alias, "v3")
	m.V4 = field.NewString(alias, "v4")
	m.V5 = field.NewString(alias, "v5")

	m.fillFieldMap()

	return &m
}

func (m *managePermission) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *managePermission) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 8)
	m.fieldMap["id"] = m.ID
	m.fieldMap["ptype"] = m.Ptype
	m.fieldMap["v0"] = m.V0
	m.fieldMap["v1"] = m.V1
	m.fieldMap["v2"] = m.V2
	m.fieldMap["v3"] = m.V3
	m.fieldMap["v4"] = m.V4
	m.fieldMap["v5"] = m.V5
}

func (m managePermission) clone(db *gorm.DB) managePermission {
	m.managePermissionDo.ReplaceDB(db)
	return m
}

type managePermissionDo struct{ gen.DO }

func (m managePermissionDo) Debug() *managePermissionDo {
	return m.withDO(m.DO.Debug())
}

func (m managePermissionDo) WithContext(ctx context.Context) *managePermissionDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m managePermissionDo) Clauses(conds ...clause.Expression) *managePermissionDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m managePermissionDo) Not(conds ...gen.Condition) *managePermissionDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m managePermissionDo) Or(conds ...gen.Condition) *managePermissionDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m managePermissionDo) Select(conds ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m managePermissionDo) Where(conds ...gen.Condition) *managePermissionDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m managePermissionDo) Order(conds ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m managePermissionDo) Distinct(cols ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m managePermissionDo) Omit(cols ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m managePermissionDo) Join(table schema.Tabler, on ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m managePermissionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m managePermissionDo) RightJoin(table schema.Tabler, on ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m managePermissionDo) Group(cols ...field.Expr) *managePermissionDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m managePermissionDo) Having(conds ...gen.Condition) *managePermissionDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m managePermissionDo) Limit(limit int) *managePermissionDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m managePermissionDo) Offset(offset int) *managePermissionDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m managePermissionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *managePermissionDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m managePermissionDo) Unscoped() *managePermissionDo {
	return m.withDO(m.DO.Unscoped())
}

func (m managePermissionDo) Create(values ...*model.ManagePermission) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m managePermissionDo) CreateInBatches(values []*model.ManagePermission, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m managePermissionDo) Save(values ...*model.ManagePermission) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m managePermissionDo) First() (*model.ManagePermission, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ManagePermission), nil
	}
}

func (m managePermissionDo) Take() (*model.ManagePermission, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ManagePermission), nil
	}
}

func (m managePermissionDo) Last() (*model.ManagePermission, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ManagePermission), nil
	}
}

func (m managePermissionDo) Find() ([]*model.ManagePermission, error) {
	result, err := m.DO.Find()
	return result.([]*model.ManagePermission), err
}

func (m managePermissionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ManagePermission, err error) {
	buf := make([]*model.ManagePermission, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m managePermissionDo) FindInBatches(result *[]*model.ManagePermission, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m managePermissionDo) Attrs(attrs ...field.AssignExpr) *managePermissionDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m managePermissionDo) Assign(attrs ...field.AssignExpr) *managePermissionDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m managePermissionDo) Joins(field field.RelationField) *managePermissionDo {
	return m.withDO(m.DO.Joins(field))
}

func (m managePermissionDo) Preload(field field.RelationField) *managePermissionDo {
	return m.withDO(m.DO.Preload(field))
}

func (m managePermissionDo) FirstOrInit() (*model.ManagePermission, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ManagePermission), nil
	}
}

func (m managePermissionDo) FirstOrCreate() (*model.ManagePermission, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ManagePermission), nil
	}
}

func (m managePermissionDo) FindByPage(offset int, limit int) (result []*model.ManagePermission, count int64, err error) {
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

func (m managePermissionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m *managePermissionDo) withDO(do gen.Dao) *managePermissionDo {
	m.DO = *do.(*gen.DO)
	return m
}
