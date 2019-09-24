package services

import (
	"code_jzggxx.com/ouer/admin/app/database"
	"code_jzggxx.com/ouer/admin/app/model"
	"fmt"
	"strconv"
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

/**
分页方法
*/
func Pagination(count int, pageNum int, nowPage int, url string) string {
	var str string
	if count == 0 {
		str = ""
		return str
	}
	//计算总页数
	var totalPage = (count / pageNum) + 1
	fmt.Println("totalpage is " + strconv.Itoa(totalPage))
	if totalPage == 1 {
		str = ""
		return str
	}
	var arrNum = totalPage
	var returnString string
	returnString = "<ul class='pagination pagination-sm inline'><li><a href='" + url + "'>上一页</a></li>"
	for i := 1; i <= arrNum; i++ {
		var child = ""
		if nowPage == i {
			child = "<li><a href='" + url + "'>" + strconv.Itoa(i) + "</a></li>"
		} else {
			child = "<li><a href='" + url + "'>" + strconv.Itoa(i) + "</a></li>"
		}
		returnString = returnString + child
	}
	returnString = returnString + "<li><a href='" + url + "'>下一页</a></li></ul>"
	return returnString
}
