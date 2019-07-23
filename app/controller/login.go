package controller

import (
	"code_jzggxx.com/ouer/admin/app/database"
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	//查询admin_users表查询用户
	database.MasterDB.Table("admin_user").Where("phone = ? AND password = ?", Phone, PassWord).Find(&AdminUsers)
	jsonUsers, err := json.Marshal(AdminUsers[0])
	if err != nil {
		log.Fatal("转json失败")
	}
	//session操作
	session := sessions.Default(c)
	session.Set("user_info", string(jsonUsers))
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"res":  true,
		"data": AdminUsers,
	})
}
