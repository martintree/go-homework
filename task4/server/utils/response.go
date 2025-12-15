package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode = 0
	ErrorCode   = 1
)

// Response 统一 API 响应格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // omitempty：当 data 为 nil 时不输出字段
}

// Success 返回成功响应
func Success(c *gin.Context, data interface{}) {
	resp := Response{
		Code:    SuccessCode,
		Message: "success",
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

// SuccessWithMsg 自定义成功消息
func SuccessWithMsg(c *gin.Context, msg string, data interface{}) {
	resp := Response{
		Code:    SuccessCode,
		Message: msg,
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}

// Fail 返回业务失败响应（HTTP 200 + code != 0）
func Fail(c *gin.Context, code int, message string) {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	c.JSON(http.StatusOK, resp)
}

// Error 返回 HTTP 错误响应（如 400, 401, 500）
func Error(c *gin.Context, httpStatus int, code int, message string) {
	resp := Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
	c.JSON(httpStatus, resp)
}
