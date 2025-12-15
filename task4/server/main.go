package main

import (
	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/middleware"
	"metanode.com/homework/server/routes"
	"metanode.com/homework/server/utils"
)

func main() {
	// 初始化日志和数据库
	utils.InitLogger()
	defer utils.Logger.Sync() // 刷新缓冲区

	r := gin.New() // 使用 gin.New() 而不是 gin.Default() 以完全控制中间件

	// 中间件顺序很重要！
	r.Use(middleware.LoggerMiddleware()) // 请求日志
	r.Use(middleware.ErrorHandler())     // 全局错误处理
	r.Use(gin.Recovery())                // panic 恢复

	// 设置路由
	routes.SetupRoutes(r)

	r.Run(":8080")
}
