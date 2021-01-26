package logic

import (
	"app/model"
	"sort"
	"strconv"
	"strings"
)

func MenuLists(groupid uint8) (re []model.ListsRes) {
	items := (&model.Menu{}).SelectMenuOrderByPidASC()

	var menuArr []string
	if groupid > 0 {
		groupMenuInfo := (&model.AuthGroupMenu{}).SelectGroupMenu(groupid)
		menuArr = strings.Split(groupMenuInfo.Menu, ",")
	}

	var listsRes []model.ListsRes
	for _, v := range items {
		listsRes = append(listsRes, model.ListsRes{
			ID:         v.ID,
			Title:      v.Title,
			Index:      v.Index,
			Icon:       v.Icon,
			Breadcrumb: v.Breadcrumb,
			Real:       v.Real,
			Show:       v.Show,
			Pid:        v.Pid,
			Sort:       v.Sort,
			IsShow:     InArray(menuArr, strconv.Itoa(int(v.ID))),
		})
	}
	for _, v := range listsRes {
		if v.Pid != 0 {
			break
		}

		v.Child = (&model.Menu{ID: v.ID}).AppendChild(listsRes)
		re = append(re, v)
	}

	return re
}

func MenuInfo(user *model.AuthUser) (re []model.Router) {
	menuInfo := (&model.Menu{}).All()

	var groupIDArr []string
	for _, groupID := range user.GroupID {
		groupIDArr = append(groupIDArr, strconv.Itoa(int(groupID)))
	}

	menuKV := map[string]uint{}

	res := (&model.Menu{}).SelectForGroupID(groupIDArr)
	for _, groupRes := range res {
		for _, m := range strings.Split(groupRes.Menu, ",") {
			menuKV[m] = 1
		}
	}

	var mergeMenu []int
	for gid := range menuKV {
		if g, err := strconv.Atoi(gid); err == nil {
			mergeMenu = append(mergeMenu, g)
		}
	}

	var mergeMenuStr []string
	sort.Ints(mergeMenu)
	for _, gid := range mergeMenu {
		mergeMenuStr = append(mergeMenuStr, strconv.Itoa(gid))
	}
	m := &model.AuthGroupMenu{}
	m.Menu = strings.Join(mergeMenuStr, ",")

	for _, menu := range menuInfo {
		if menu.Pid == 0 {
			r := menuConv(m, menu, user)
			r.Children = menuGetChild(m, menu, menuInfo, user)
			for _, mm := range r.Children {
				if mm.Meta.Show {
					r.Meta.Collapse = true
				}

				if mm.Meta.Has && r.Name != "后台中心" {
					r.Url = ""
				}
			}

			re = append(re, r)
		}
	}

	return re
}

func menuConv(m *model.AuthGroupMenu, menu model.Menu, user *model.AuthUser) (r model.Router) {
	show := InArray(append(strings.Split(m.Menu, ","), "1"), strconv.Itoa(int(menu.ID)))
	has := InArray(append(strings.Split(m.Menu, ","), "1", "2", "7"), strconv.Itoa(int(menu.ID)))
	if user.IsSuper {
		has = true
		show = true
	}

	r = model.Router{
		Name: menu.Title,
		Path: MenuVuePath(menu.Index),
		// Url:        MenuVueUrl(InArray(strings.Split(m.Menu, ","), strconv.Itoa(int(menu.ID))), menu.Index),
		Url:      MenuVueUrl(true, menu.Index),
		Icon:     menu.Icon,
		Children: []model.Router{},
	}
	r.Meta.Breadcrumb = menu.Breadcrumb == 1
	r.Meta.Real = menu.Real == 1
	r.Meta.Show = menu.Show == 1 && show
	r.Meta.Has = has
	r.Meta.Collapse = false

	return
}

func menuGetChild(m *model.AuthGroupMenu, menu model.Menu, menus []model.Menu, user *model.AuthUser) []model.Router {
	re := make([]model.Router, 0)
	for _, v := range menus {
		if v.Pid == uint8(menu.ID) {
			re = append(re, menuConv(m, v, user))
		}
	}

	return re
}

func MenuVueUrl(show bool, url string) string {
	if !show {
		return ""
	}

	if url == "main" {
		return "pages/main/" + url + ".vue"
	}

	if strings.HasPrefix(url, "/") {
		return "pages" + url + ".vue"
	}

	return "pages/" + url + ".vue"
}

func MenuVuePath(path string) string {
	if strings.HasPrefix(path, "/") {
		if !strings.HasPrefix(path, "/main") {
			return "/main" + path
		}
	} else {
		if path == "main" {
			return "/" + path + "/main"
		} else if !strings.HasPrefix(path, "/main/") {
			return "/main/" + path
		}
	}

	return path
}
