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
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// 禁用控制台颜色
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("app/storage/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)


	//初始化配置
	config.Init()
	//初始化数据库链接
	database.Init()
	//静态文件
	////初始化gin
	//r := gin.Default()
	//创建无中间件路由
	r := gin.New()
	// 使用 Logger 中间件
	r.Use(gin.Logger())

	// 使用 Recovery 中间件
	r.Use(gin.Recovery())

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
