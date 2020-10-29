package model

import (
	"app/web/business/manageBusiness"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// AuthGroupMenu 角色菜单对应表
type AuthGroupMenu struct {
	ID        uint           `gorm:"column:id;primaryKey;" json:"id,omitempty"`
	GroupID   uint8          `gorm:"column:groupid;type:tinyint(4);not null;default:0;comment:角色Id;" json:"groupid"`
	Menu      string         `gorm:"column:menu;type:varchar(255);not null;default:'';comment:菜单;" json:"menu"`
	CreatedAt JSONTime       `gorm:"column:create_time;type:datetime(0);comment:创建时间;" json:"create_time"`
	UpdatedAt JSONTime       `gorm:"column:update_time;type:datetime(0);comment:更新时间;" json:"update_time"`
	DeletedAt gorm.DeletedAt `gorm:"type:datetime(0);index;" json:"-"`
}

var initNum uint = 3

func (*migrate) CreateAuthGroupMenu() {
	migrateData = append(migrateData, func() (string, func(db *gorm.DB) error) {
		return "CreateAuthGroupMenu", func(db *gorm.DB) error {
			db.Create([]AuthGroupMenu{
				{
					GroupID: 1,
					Menu:    "1,2,3,4,5,6,7,8,9,10,11",
				},
				{
					GroupID: 2,
					Menu:    "1",
				},
			})
			return nil
		}
	})
}

func (m *AuthGroupMenu) Update() error {
	hasInfo := &AuthGroupMenu{}
	db.Where("groupid = ?", m.GroupID).First(hasInfo)
	if hasInfo.ID == 0 {
		if res := db.Create(m); res.RowsAffected == 0 {
			return errors.New("服务繁忙,请重试.")
		}
	} else {
		if res := db.Model(&m).Select("update_time", "menu").Where("id = ?", hasInfo.ID).Updates(m); res.RowsAffected == 0 {
			return errors.New("服务繁忙,请重试.")
		}
	}

	return nil
}

type Router struct {
	Name       string   `json:"name"`
	Path       string   `json:"path"`
	Url        string   `json:"url"`
	Icon       string   `json:"icon"`
	Breadcrumb bool     `json:"breadcrumb"`
	Real       bool     `json:"real"`
	Show       bool     `json:"show"`
	Collapse   bool     `json:"collapse"`
	Children   []Router `json:"children"`
}

func (m *AuthGroupMenu) GroupMenu(user *AuthUser) (res []Router) {
	menuInfo := (&Menu{}).All()

	db.Where("groupid = ?", m.GroupID).Find(&m)
	menuArr := strings.Split(m.Menu, ",")
	push := func(res []Router, v Menu, child []Router, collapse bool) []Router {
		res = append(res, Router{
			Name:       v.Title,
			Path:       m.VuePath(v.Index),
			Url:        m.VueUrl(v.Show == 1 && collapse, v.Index),
			Icon:       v.Icon,
			Breadcrumb: v.Breadcrumb == 1,
			Real:       v.Real == 1,
			Show:       v.Show == 1,
			Collapse:   collapse,
			Children:   child,
		})
		return res
	}
	// 公共菜单栏
	for _, initRouter := range menuInfo {
		if initRouter.ID <= initNum && initRouter.Pid == 0 {
			child, collapse := (&AuthGroupMenu{}).AppendChildRen(initRouter, menuInfo)
			res = push(res, initRouter, child, collapse)
		}
	}
	if user.IsSuper { // 超级拥有全部菜单栏
		for _, sysMenu := range menuInfo {
			if sysMenu.ID > initNum && sysMenu.Pid == 0 {
				child, collapse := (&AuthGroupMenu{}).AppendChildRen(sysMenu, menuInfo)
				res = push(res, sysMenu, child, collapse)
			}
		}
	} else { // 根据用户组获取设置的菜单栏
		for _, sysMenu := range menuInfo {
			if sysMenu.ID > initNum && sysMenu.Pid == 0 && manageBusiness.InArray(menuArr, strconv.Itoa(int(sysMenu.ID))) {
				child, collapse := m.AppendChildRen(sysMenu, menuInfo)
				res = push(res, sysMenu, child, collapse)
			}
		}
	}

	return res
}

func (m *AuthGroupMenu) AppendChildRen(currentMenu Menu, menuMap []Menu) (res []Router, collapse bool) {
	push := func(res []Router, v Menu) []Router {
		res = append(res, Router{
			Name:       v.Title,
			Path:       m.VuePath(v.Index),
			Url:        m.VueUrl(false, v.Index),
			Icon:       v.Icon,
			Breadcrumb: v.Breadcrumb == 1,
			Real:       v.Real == 1,
			Show:       v.Show == 1,
		})
		return res
	}

	for _, v := range menuMap {
		if currentMenu.ID <= initNum && v.ID <= initNum && currentMenu.ID == uint(v.Pid) { // 公共
			res = push(res, v)
			if v.Show == 1 {
				collapse = true
			}
		} else if currentMenu.ID > initNum && v.ID > initNum && currentMenu.ID == uint(v.Pid) {
			if m.ID > 0 { // 代表不是超级管理员
				menuRule := strings.Split(m.Menu, ",")
				if manageBusiness.InArray(menuRule, strconv.Itoa(int(v.ID))) {
					res = push(res, v)
				}
			} else {
				res = push(res, v)
			}

			if v.Show == 1 {
				collapse = true
			}
		}
	}

	return res, collapse
}

func (m *AuthGroupMenu) VueUrl(show bool, url string) string {
	if show {
		return ""
	}
	return "pages/main/" + url + ".vue"
}

func (m *AuthGroupMenu) VuePath(path string) string {
	if strings.HasPrefix(path, "/") {
		if !strings.HasPrefix(path, "/main") {
			return "/main" + path
		}
	} else {
		if path == "main" {
			return "/" + path
		} else if !strings.HasPrefix(path, "/main/") {
			return "/main/" + path
		}
	}

	return path
}
