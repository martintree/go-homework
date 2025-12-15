package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"metanode.com/homework/server/utils"
)

type Users struct {
	gorm.Model
	Username string `gorm:"size:100;uniqueIndex;not null" json:"username"`
	Email    string `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"` // 不返回给前端
}

func (u *Users) Register(tx *gorm.DB) *utils.AppError {
	if len(u.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return utils.WrapError(utils.ErrInternalServer, errors.New("failed to hash password"))
		}
		u.Password = string(hashedPassword)

		if err := tx.Create(&u).Error; err != nil {
			return utils.WrapError(utils.ErrDatabase, err)
		}

		return nil
	}
	return utils.WrapError(utils.ErrInvalidRequest, errors.New("password can not be empty"))
}

func (u *Users) Login(tx *gorm.DB) (string, *utils.AppError) {
	var storedUser Users
	if err := tx.Where("username = ?", u.Username).First(&storedUser).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", utils.ErrInvalidCredentials
		} else {
			return "", utils.WrapError(utils.ErrDatabase, err)
		}
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
		return "", utils.ErrInvalidCredentials
	}

	// 生成 JWT
	token, err := utils.GenerateToken(storedUser.ID)
	if err != nil {
		return "", utils.ErrInternalServer
	}

	return token, nil
}
