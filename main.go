package main

import (
	"github.com/gin-gonic/gin"
	"last-homework/api"
	"last-homework/dao/mysql"
	"last-homework/middleware"
	"last-homework/tool"
	"log"
)

func main() {
	//加载配置
	if err := tool.Viper(); err != nil {
		log.Fatal("viper出错:", err)
		return
	}
	//初始化日志
	if err := tool.Logger(); err != nil {
		log.Fatal("zap出错:", err)
		return
	}
	tool.SugaredDebug("zap logger初始化...")

	if err := mysql.Mysql(); err != nil {
		tool.SugaredFatalf("mysql init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("mysql 初始化...")

	if err := tool.Redis(); err != nil {
		tool.SugaredFatalf("redis init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("redis 初始化...")

	if err := tool.InitSnowflake(); err != nil {
		tool.SugaredFatalf("snowflake init error: %s", err.Error())
		return
	}
	tool.SugaredDebug("snowflake 初始化...")

	//路由
	URL()
}

func URL() {
	//注册路由
	r := gin.New()

	r.Use(middleware.GinLogger, middleware.GinZapRecovery(true))
	r.Use(middleware.Cors())

	r.GET("/", func(c *gin.Context) {
		tool.ResponseSuccess(c, gin.H{"msg": "成功"})
	})

	u := r.Group("/user")
	{
		u.POST("register", api.Register)
		u.POST("login", api.Login)
	}

	//websocket
	room := r.Group("/room")
	room.Use(middleware.Jwt())
	{
		room.GET("/create", api.CreateRoom)
		room.GET("/enter", api.EnterRoom)
		room.GET("/search", api.Search)

	}

	r.Run(tool.GetViper().GetString("app.port"))
}
