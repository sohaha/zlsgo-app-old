// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameManagePermission = "dt_manage_permission"

// ManagePermission mapped from table <dt_manage_permission>
type ManagePermission struct {
	ID    int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Ptype string `gorm:"column:ptype;default:NULL" json:"ptype"`
	V0    string `gorm:"column:v0;default:NULL" json:"v0"`
	V1    string `gorm:"column:v1;default:NULL" json:"v1"`
	V2    string `gorm:"column:v2;default:NULL" json:"v2"`
	V3    string `gorm:"column:v3;default:NULL" json:"v3"`
	V4    string `gorm:"column:v4;default:NULL" json:"v4"`
	V5    string `gorm:"column:v5;default:NULL" json:"v5"`
}

// TableName ManagePermission's table name
func (*ManagePermission) TableName() string {
	return TableNameManagePermission
}