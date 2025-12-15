package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metanode.com/homework/server/utils"
)

// LoggerMiddleware 记录每个请求
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// 处理请求
		c.Next()

		// 记录请求日志
		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()

		fields := []zap.Field{
			zap.Int("status", statusCode),
			zap.String("method", method),
			zap.String("path", path),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		}

		// 根据状态码选择日志级别
		if statusCode >= 500 {
			utils.Logger.Error("服务器错误", fields...)
		} else if statusCode >= 400 {
			utils.Logger.Warn("客户端错误", fields...)
		} else {
			utils.Logger.Info("请求完成", fields...)
		}
	}
}
