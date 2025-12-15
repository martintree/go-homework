package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"metanode.com/homework/server/utils"
)

func isNotNilAppError(appErr *utils.AppError) bool {
	return appErr != nil && appErr != (*utils.AppError)(nil)
}

// ErrorHandler 全局错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 如果已经有响应，跳过
		if c.Writer.Written() {
			return
		}

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// 类型断言为 AppError
			var appErr *utils.AppError
			if errors.As(err, &appErr) && isNotNilAppError(appErr) {
				// 记录详细错误日志（仅内部）
				if appErr.Err != nil {
					utils.Logger.Error("应用错误",
						zap.Int("http_code", appErr.Code),
						zap.String("message", appErr.Message),
						zap.Error(appErr.Err),
						zap.String("path", c.Request.URL.Path),
						zap.String("method", c.Request.Method),
					)
				}

				utils.Fail(c, appErr.Code, appErr.Message)
				return
			}

			// 未知错误 → 500
			utils.Logger.Error("未处理的错误",
				zap.Error(err),
				zap.String("path", c.Request.URL.Path),
				zap.String("method", c.Request.Method),
			)
			utils.Fail(c, utils.ErrInvalidRequest.Code, utils.ErrInvalidRequest.Message)
		}
	}
}
