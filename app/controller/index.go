package controller

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq"
	"log"
	"net/http"
)

/**
后台首页
*/
func Index(c *gin.Context) {
	session := sessions.Default(c)
	userInfo := session.Get("user_info")
	var users AdminUser
	err := json.Unmarshal([]byte(userInfo.(string)), &users)
	if err != nil {
		log.Fatal("json解析失败")
	}
	nickName := gojsonq.New().JSONString(userInfo.(string)).Find("nick_name") //取参数方式一
	nickName = users.NickName                                                 //取参数方式二
	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"name":      users.Name,
		"nick_name": nickName,
	})
}
