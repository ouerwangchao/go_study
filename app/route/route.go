package route

import (
	"code_jzggxx.com/ouer/admin/app/controller"
	"code_jzggxx.com/ouer/admin/app/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(r *gin.Engine) {
	//404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.LoadHTMLGlob("app/view/*/*")
	//Api相关路由
	a := r.Group("/api")
	a.GET("/getUser", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "api路由",
		})
	})
	//layout
	//r.HTMLRender = createMyRender()
	//Home相关路由
	h := r.Group("/home")
	h.GET("/login", controller.Login)               //登录页面
	h.POST("/login/action", controller.LoginAction) //登录操作
	h.Use(middleware.CheckLogin)
	{
		h.GET("/index", controller.Index) //后台首页
	}
}
