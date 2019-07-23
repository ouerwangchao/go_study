package main

import (
	"code_jzggxx.com/ouer/admin/app/config"
	"code_jzggxx.com/ouer/admin/app/database"
	"code_jzggxx.com/ouer/admin/app/route"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	//初始化配置
	config.Init()
	//初始化数据库链接
	database.Init()
	//静态文件
	//初始化gin
	r := gin.Default()

	//session相关
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//静态文件
	r.Static("app/public", "./app/public")
	r.StaticFS("/static", http.Dir("./app/public"))
	r.StaticFile("/favicon.ico", "favicon.ico")
	route.SetupRouter(r)
	ser := &http.Server{
		Addr:         ":" + viper.GetString("app.port"),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	gracehttp.Serve(ser)

}
