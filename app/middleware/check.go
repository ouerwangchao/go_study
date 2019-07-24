package middleware

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/**
检查是否处于登录状态
*/
func CheckLogin(c *gin.Context) {
	fmt.Println(1)
	session := sessions.Default(c)
	//获取用户登录信息
	userInfo := session.Get("user_info")
	fmt.Println(userInfo)
	c.Next()
}
