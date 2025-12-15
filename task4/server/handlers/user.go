package handlers

import (
	"github.com/gin-gonic/gin"
	"metanode.com/homework/server/db"
	"metanode.com/homework/server/dto"
	"metanode.com/homework/server/utils"
)

func RegisterUser(c *gin.Context) {
	var toAddUser dto.UserCreateRequest
	if err := c.ShouldBindJSON(&toAddUser); err != nil {
		c.Error(err)
		return
	}

	//dto转model
	user := dto.ToCreateUserModel(&toAddUser)
	if err := user.Register(db.GetDB()); err != nil {
		c.Error(err)
		return
	}
	utils.Success(c, user)
}

// 登录
func Login(c *gin.Context) {
	var loginUser dto.UserLoginRequest
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		//utils.Fail(c, utils.ErrInvalidRequest.Code, utils.ErrInvalidRequest.Message)
		c.Error(err)
		return
	}

	//dto转model
	user := dto.ToLoginUserModel(&loginUser)
	token, err := user.Login(db.GetDB())
	if err != nil {
		c.Error(err)
		return
	}
	utils.Success(c, token)
}
