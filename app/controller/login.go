package controller

import (
	"code_jzggxx.com/ouer/admin/app/database"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

/**
登录页面
*/
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.html", gin.H{})
}

/**
接收admin_user表数据
*/
type AdminUser struct {
	Id        int       `json:"id"`
	RoleId    int       `json:"role_id"`
	NickName  string    `json:"nick_name"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Conf      string    `json:"conf"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

/**
登录操作
*/
func LoginAction(c *gin.Context) {
	var AdminUsers []AdminUser
	Phone := c.PostForm("phone")
	PassWord := md5V(c.PostForm("password"))
	//查询admin_user表查询用户
	database.MasterDB.Table("admin_user").Where("phone = ? AND password = ?", Phone, PassWord).Find(&AdminUsers)
	fmt.Println(AdminUsers)
	if len(AdminUsers) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"res":  false,
			"data": "账号/密码错误",
		})
		return
	}
	jsonUsers, err := json.Marshal(AdminUsers[0])
	if err != nil {
		log.Fatal("转json失败")
	}
	//session操作
	session := sessions.Default(c)
	//session.Set("user_info", string(jsonUsers))
	session.Set(viper.GetString("enums.admin.session_user_info"), string(jsonUsers))
	session.Save()
	fmt.Println(session.Get("user_info"))

	c.JSON(http.StatusOK, gin.H{
		"res":  true,
		"data": AdminUsers,
	})
	return
}
