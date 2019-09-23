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
	var AdminUser []AdminUser
	if Phone == "" {
		database.MasterDB.Table("admin_user").Find(&AdminUser)
	} else {
		database.MasterDB.Table("admin_user").Where("phone = ? ", Phone).Find(&AdminUser)
	}
	var Users []ShowAdminUser
	for i, value := range AdminUser {
		var userArr ShowAdminUser
		userArr.Id = i + 1
		userArr.Phone = value.Phone
		userArr.Name = value.Name
		userArr.NickName = value.NickName
		var dateTime = value.CreatedAt
		userArr.CreatedAt = dateTime.Format("2006-01-02 15:04:05")
		Users = append(Users, userArr)
	}
	menu := services.GetAdminMenu()
	fmt.Println(menu);
	c.HTML(http.StatusOK, "users/index.html", gin.H{
		"data": Users,
		"menu": menu,
	})
}

type ShowAdminUser struct {
	Id        int
	NickName  string
	Name      string
	Phone     string
	CreatedAt string
}

type AddAdminUser struct {
	NickName string `json:"nick_name"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func EditUser(c *gin.Context) {
	var AdminUsers AddAdminUser
	AdminUsers.Phone = c.PostForm("phone")
	AdminUsers.Name = c.PostForm("name")
	AdminUsers.NickName = c.PostForm("nick_name")
	AdminUsers.Password = md5V("123456")
	fmt.Println(AdminUsers)
	if err := database.MasterDB.Table("admin_user").Create(AdminUsers).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"res":  false,
			"data": "新增失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"res":  true,
		"data": "操作成功",
	})
	return
}
