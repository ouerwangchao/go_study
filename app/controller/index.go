package controller

import (
	service "code_jzggxx.com/ouer/admin/app/services"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/thedevsaddam/gojsonq"
	"log"
	"net/http"
)

/**
后台首页
*/
func Index(c *gin.Context) {
	session := sessions.Default(c)
	userInfo := session.Get(viper.GetString("enums.admin.session_user_info"))
	if userInfo == nil {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
	var users AdminUser
	err := json.Unmarshal([]byte(userInfo.(string)), &users)
	if err != nil {
		log.Fatal("json解析失败")
	}
	nickName := gojsonq.New().JSONString(userInfo.(string)).Find("nick_name") //取参数方式一
	nickName = users.NickName                                                 //取参数方式二

	//获取导航菜单栏数据
	menu := service.GetAdminMenu()
	fmt.Println(menu)
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"name":      users.Name,
		"nick_name": nickName,
		"menu":      menu,
	})
}
