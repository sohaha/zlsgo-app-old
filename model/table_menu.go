package model

import (
	"errors"
	"gorm.io/gorm"
)

// Menu 菜单表
type Menu struct {
	ID         uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	Title      string         `gorm:"column:title;type:varchar(255);not null;default:'';comment:菜单名称;" json:"title"`
	Index      string         `gorm:"column:index;type:varchar(255);not null;default:'';comment:路由;" json:"index"`
	Icon       string         `gorm:"column:icon;type:varchar(255);not null;default:'';comment:图标;" json:"icon"`
	Breadcrumb uint8          `gorm:"column:breadcrumb;type:tinyint(4);not null;default:0;comment:面包屑显示:0,1;" json:"breadcrumb"`
	Real       uint8          `gorm:"column:real;type:tinyint(4);not null;default:0;comment:面包屑可点击:0,1;" json:"real"`
	Show       uint8          `gorm:"column:show;type:tinyint(4);not null;default:0;comment:导航栏显示:0,1;" json:"show"`
	Pid        uint8          `gorm:"column:pid;type:tinyint(4);not null;default:0;comment:父id;" json:"pid"`
	Sort       uint8          `gorm:"column:sort;type:tinyint(4);not null;default:0;comment:排序id;" json:"sort"`
	CreatedAt  JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt  JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt  gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

func (*migrate) CreateMenu() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateMenu", func(db *gorm.DB) error {
			tx := db.Create([]Menu{
				{Title: "后台中心", Index: "main", Icon: "icon-pie-chart-", Breadcrumb: 1, Real: 1, Show: 1, Pid: 0, Sort: 1},
				{Title: "站内消息", Index: "user/logs", Icon: "icon-settings-", Breadcrumb: 0, Real: 0, Show: 0, Pid: 1, Sort: 2},
				{Title: "多端登录", Index: "user/client", Icon: "icon-globe--outline", Breadcrumb: 0, Real: 0, Show: 0, Pid: 1, Sort: 3},
				{Title: "日志查看", Index: "system/logs", Icon: "icon-alert-circle", Breadcrumb: 0, Real: 0, Show: 1, Pid: 0, Sort: 4},
				{Title: "系统设置", Index: "system", Icon: "icon-options-", Breadcrumb: 0, Real: 0, Show: 1, Pid: 0, Sort: 5},
				{Title: "程序设置", Index: "system/config", Icon: "icon-settings", Breadcrumb: 1, Real: 0, Show: 1, Pid: 5, Sort: 6},
				{Title: "用户设置", Index: "user/lists", Icon: "icon-person", Breadcrumb: 1, Real: 0, Show: 1, Pid: 5, Sort: 7},
				{Title: "角色设置", Index: "user/group", Icon: "icon-people", Breadcrumb: 1, Real: 0, Show: 1, Pid: 5, Sort: 8},
				{Title: "菜单设置", Index: "user/menu", Icon: "icon-pricetags", Breadcrumb: 1, Real: 0, Show: 1, Pid: 5, Sort: 9},
				{Title: "个人设置", Index: "user/my", Icon: "icon-person-done", Breadcrumb: 1, Real: 0, Show: 1, Pid: 5, Sort: 10},
				{Title: "权限设置", Index: "user/rules", Icon: "icon-pantone", Breadcrumb: 1, Real: 0, Show: 1, Pid: 5, Sort: 11},
			})
			return tx.Error
		}
	})
}

type ListsRes struct {
	ID         uint       `json:"id"`
	Title      string     `json:"title"`
	Index      string     `json:"index"`
	Icon       string     `json:"icon"`
	Breadcrumb uint8      `json:"breadcrumb"`
	Real       uint8      `json:"real"`
	Show       uint8      `json:"show"`
	Pid        uint8      `json:"pid"`
	Sort       uint8      `json:"sort"`
	IsShow     bool       `json:"is_show"`
	Child      []ListsRes `json:"child"`
}

func (m *Menu) SelectMenuOrderByPidASC() []Menu {
	var items []Menu
	db.Model(&m).Order("pid asc").Order("sort asc").Find(&items)

	return items
}

// AppendChild 追加子菜单
func (m *Menu) AppendChild(sourceData []ListsRes) (res []ListsRes) {
	for _, v := range sourceData {
		if uint8(m.ID) == v.Pid {
			res = append(res, v)
		}
	}

	return
}

func (m *Menu) Exist() {
	db.First(&m, m.ID)
}

func (m *Menu) PidExist() {
	db.Where("pid = ?", m.Pid).First(&m)
}

func (m *Menu) Create() error {
	if res := db.Select("title", "index", "icon", "breadcrumb", "real", "show", "pid", "sort").Create(&m); res.RowsAffected == 0 {
		return errors.New("服务繁忙，请重试")
	}

	return nil
}

func (m *Menu) Delete() error {
	if res := db.Delete(&m); res.RowsAffected == 0 {
		return errors.New("服务繁忙，请重试")
	}

	return nil
}

func (m *Menu) Update() error {
	if res := db.Model(&m).Select("update_time", "title", "index", "icon", "breadcrumb", "real", "show").Where("id = ?", m.ID).Updates(m); res.RowsAffected == 0 {
		return errors.New("服务繁忙，请重试")
	}

	return nil
}

type PostSortSt []struct {
	ID    int `json:"id"`
	Child []struct {
		ID int `json:"id"`
	} `json:"child,omitempty"`
}

func (m *Menu) MenuSort(data PostSortSt) error {
	// TODO
	// 这里的更新可以优化为同时一个sql更新多条不一样的数据
	tx := db.Begin()
	i := 1
	for _, v := range data {
		uRes := &Menu{ID: uint(v.ID), Sort: uint8(i)}
		if res := tx.Model(&m).Select("update_time", "sort").Where("id = ?", uRes.ID).Updates(uRes); res.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("服务繁忙，请重试")
		}
		i++
		if len(v.Child) > 0 {
			for _, vv := range v.Child {
				uRes2 := &Menu{ID: uint(vv.ID), Sort: uint8(i), Pid: uint8(v.ID)}
				if res := tx.Model(&m).Select("update_time", "sort", "pid").Where("id = ?", uRes2.ID).Updates(uRes2); res.RowsAffected == 0 {
					tx.Rollback()
					return errors.New("服务繁忙，请重试")
				}
				i++
			}
		}
	}
	tx.Commit()

	return nil
}

func (m *Menu) All() (res []Menu) {
	db.Model(m).Order("sort asc").Find(&res)
	return
}

func (m *Menu) SelectForGroupID(groupIDArr []string) []AuthGroupMenu {
	var res []AuthGroupMenu
	db.Where("groupid IN ?", groupIDArr).Find(&res)

	return res
}
