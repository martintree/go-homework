package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/db"
	"metanode.com/homework/server/dto"
)

func RegisterUser(c *gin.Context) {
	var toAddUser dto.UserCreateRequest
	if err := c.ShouldBindJSON(&toAddUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "invalid params"})
		return
	}

	//dto转model
	user := dto.ToCreateUserModel(&toAddUser)
	if err := user.Register(db.GetDB()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user, "error": ""})
}

// 登录
func Login(c *gin.Context) {
	var loginUser dto.UserLoginRequest
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": "invalid params"})
		return
	}

	//dto转model
	user := dto.ToLoginUserModel(&loginUser)
	token, err := user.Login(db.GetDB())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": token, "error": ""})
}
