package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

/**
检查是否处于登录状态
*/
func CheckLogin(c *gin.Context) {
	fmt.Println(11112222211)
	session := sessions.Default(c)
	//获取用户登录信息
	userInfo := session.Get(viper.GetString("enums.admin.session_user_info"))
	fmt.Println(userInfo)
	if userInfo == nil {
		fmt.Println("No login")
		c.Redirect(http.StatusMovedPermanently, "/home/login")
	}
	c.Next()
}
