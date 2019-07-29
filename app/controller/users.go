package controller

import (
	"code_jzggxx.com/ouer/admin/app/database"
	"code_jzggxx.com/ouer/admin/app/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
后台管理员页面
*/
func Users(c *gin.Context) {
	Phone := c.DefaultQuery("phone", "")
	fmt.Println(Phone)
	var AdminUser []AdminUser
	if Phone == "" {
		database.MasterDB.Table("admin_user").Find(&AdminUser)
	} else {
		database.MasterDB.Table("admin_user").Where("phone = ? ", Phone).Find(&AdminUser)
	}
	menu := services.GetAdminMenu()
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"data": AdminUser,
		"menu": menu,
	})
}
