package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/utils"
)

// AuthMiddleware JWT 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"data": "", "error": "authorization header required"})
			return
		}

		// 格式: Bearer <token>
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"data": "", "error": "invalid Authorization header format"})
			return
		}

		tokenStr := headerParts[1]
		claims, err := utils.ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"data": "", "error": err.Error()})
			return
		}

		// 将用户ID存入上下文
		c.Set("userId", claims.UserID)
		c.Next()
	}
}
