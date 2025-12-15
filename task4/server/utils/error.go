package utils

import (
	"net/http"
)

// AppError 应用级错误结构
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"` // 内部错误，不暴露给客户端
}

// Error 实现 error 接口
func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

// 错误构造函数
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// 预定义常见错误
var (
	ErrInvalidRequest     = NewAppError(http.StatusBadRequest, "请求参数无效", nil)
	ErrUnauthorized       = NewAppError(http.StatusUnauthorized, "未授权访问", nil)
	ErrForbidden          = NewAppError(http.StatusForbidden, "无权限操作", nil)
	ErrNotFound           = NewAppError(http.StatusNotFound, "资源不存在", nil)
	ErrConflict           = NewAppError(http.StatusConflict, "资源已存在", nil)
	ErrInternalServer     = NewAppError(http.StatusInternalServerError, "服务器内部错误", nil)
	ErrDatabase           = NewAppError(http.StatusInternalServerError, "数据库操作失败", nil)
	ErrInvalidCredentials = NewAppError(http.StatusUnauthorized, "用户名或密码错误", nil)
	ErrTokenExpired       = NewAppError(http.StatusUnauthorized, "登录已过期，请重新登录", nil)
	ErrInvalidToken       = NewAppError(http.StatusUnauthorized, "无效的认证凭证", nil)
)

// WrapError 包装底层错误
func WrapError(baseErr *AppError, err error) *AppError {
	return &AppError{
		Code:    baseErr.Code,
		Message: baseErr.Message,
		Err:     err,
	}
}
