package routes

import (
	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/handlers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		//用户
		api.POST("/users/register", handlers.RegisterUser)
		api.POST("/users/login", handlers.Login)
		//文章
		api.GET("users/:userId/posts", handlers.GetPosts) // 获取文章列表
		api.GET("/posts/:id", handlers.GetPost)           // 获取单篇文章
		api.POST("/posts", handlers.AddPost)              // 创建文章
		api.PUT("/posts", handlers.UpdatePost)            // 更新文章
		api.DELETE("/posts/:id", handlers.DeletePost)     // 删除文章
		//评论
		api.POST("/comments", handlers.AddComment)              // 创建评论
		api.DELETE("/comments/:postId", handlers.DeleteComment) // 删除评论

	}
}
