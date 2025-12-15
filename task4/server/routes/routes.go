package routes

import (
	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/handlers"
	"metanode.com/homework/server/middleware"
)

func SetupRoutes(r *gin.Engine) {

	//不需要认证的接口
	unProtected := r.Group("/api/v1")
	{
		//用户
		unProtected.POST("/users/register", handlers.RegisterUser)
		unProtected.POST("/users/login", handlers.Login)
	}
	//需要认证的接口
	protected := r.Group("/api/v1/auth")
	// 添加认证中间件
	protected.Use(middleware.AuthMiddleware())
	{
		//文章
		protected.GET("users/posts", handlers.GetPosts)     // 获取文章列表
		protected.GET("/posts/:id", handlers.GetPost)       // 获取单篇文章
		protected.POST("/posts", handlers.AddPost)          // 创建文章
		protected.PUT("/posts", handlers.UpdatePost)        // 更新文章
		protected.DELETE("/posts/:id", handlers.DeletePost) // 删除文章
		//评论
		protected.POST("/comments", handlers.AddComment)                       // 创建评论
		protected.DELETE("/comments/post/:postId", handlers.DeletePostComment) // 删除某个文章下的所有评论
		protected.DELETE("/comments/:id", handlers.DeleteComment)              // 删除评论

	}
}
