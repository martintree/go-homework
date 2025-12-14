package dto

import "metanode.com/homework/server/models"

type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=1,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required,min=1,max=100"`
	Password string `json:"password" binding:"required,min=6,max=100"`
}

func ToCreateUserModel(req *UserCreateRequest) *models.Users {
	return &models.Users{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToLoginUserModel(req *UserLoginRequest) *models.Users {
	return &models.Users{
		Username: req.Username,
		Password: req.Password,
	}
}
