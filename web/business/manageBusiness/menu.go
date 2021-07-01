package manageBusiness

import (
	"app/logic"
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
			IsShow:     logic.InArray(menuArr, strconv.Itoa(int(v.ID))),
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

	mergeMenuStr := []string{"1", "2", "8"}
	sort.Ints(mergeMenu)
	for _, gid := range mergeMenu {
		mergeMenuStr = append(mergeMenuStr, strconv.Itoa(gid))
	}
	m := &model.AuthGroupMenu{}
	m.Menu = strings.Join(mergeMenuStr, ",")

	for _, menu := range menuInfo {
		if menu.Pid == 0 {
			r := menuConv(menu)
			r.Children = menuGetChild(m, menu, menuInfo)

			appendFlag := false
			if len(r.Children) > 0 {
				appendFlag = true
			}

			if !appendFlag {
				for _, groupId := range strings.Split(m.Menu, ",") {
					if strconv.Itoa(int(menu.ID)) == groupId {
						appendFlag = true
						break
					}
				}
			}

			if appendFlag {
				re = append(re, r)
			}
		}
	}

	for topMenuKey, topMenu := range re {
		for _, subMenu := range topMenu.Children {
			if subMenu.Meta.Show == true {
				re[topMenuKey].Url = ""
				re[topMenuKey].Meta.Collapse = true
			}
		}
	}

	return re
}

func menuConv(menu model.Menu) (r model.Router) {
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
	r.Meta.Show = menu.Show == 1
	r.Meta.Collapse = false

	return
}

func menuGetChild(m *model.AuthGroupMenu, menu model.Menu, menus []model.Menu) []model.Router {
	re := make([]model.Router, 0)
	for _, v := range menus {
		if v.Pid == uint8(menu.ID) {
			for _, groupId := range strings.Split(m.Menu, ",") {
				if strconv.Itoa(int(v.ID)) == groupId {
					re = append(re, menuConv(v))
					break
				}
			}
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
