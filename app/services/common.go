package services

import (
	"code_jzggxx.com/ouer/admin/app/database"
	"code_jzggxx.com/ouer/admin/app/model"
)

type Child struct {
	Title string
	Icon  string
	Uri   string
	Id    int
}

type Menu struct {
	Title string
	Id    int
	Class string
	Style string
	Icon  string
	Uri   string
	Child []Child
}

/**
获取后台导航栏
*/
func GetAdminMenu() []Menu {
	var AdminMenus []Menu
	//AdminMenus := make([]Menu, 0)
	//查询父级菜单
	var ParentMenu []model.AdminMenu
	var ChildMenu []model.AdminMenu
	display := 0
	database.MasterDB.Where("display = ? AND parent_id = 0", display).Find(&ParentMenu)
	database.MasterDB.Where("display = ? AND parent_id != 0", display).Find(&ChildMenu)
	for _, val := range ParentMenu {
		var adminMenu Menu
		adminMenu.Id = val.Id
		adminMenu.Title = val.Title
		adminMenu.Uri = val.Uri
		adminMenu.Icon = val.Icon
		adminMenu.Class = "treeview"
		adminMenu.Style = "display: none;"
		for _, value := range ChildMenu {
			var childMenu Child
			if value.ParentId == val.Id {
				childMenu.Uri = value.Uri
				childMenu.Title = value.Title
				childMenu.Icon = value.Icon
				childMenu.Id = value.Id
				adminMenu.Child = append(adminMenu.Child, childMenu)
			}
		}
		AdminMenus = append(AdminMenus, adminMenu)
	}
	return AdminMenus
}
