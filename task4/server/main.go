package main

import (
	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/routes"
)

func main() {
	r := gin.Default()
	// 设置路由
	routes.SetupRoutes(r)

	r.Run(":8080")
}
