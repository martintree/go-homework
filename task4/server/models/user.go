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

func (u *Users) Register(tx *gorm.DB) error {
	if len(u.Password) > 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash password")
		}
		u.Password = string(hashedPassword)

		if err := tx.Create(&u).Error; err != nil {
			return errors.New("failed to create user")
		}

		return nil
	}
	return errors.New("password can not be empty")
}

func (u *Users) Login(tx *gorm.DB) (string, error) {
	var storedUser Users
	if err := tx.Where("username = ?", u.Username).First(&storedUser).Error; err != nil {
		return "", errors.New("invalid username or password")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
		return "", errors.New("invalid username or password")
	}

	// 生成 JWT
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"id":       storedUser.ID,
	// 	"username": storedUser.Username,
	// 	"exp":      time.Now().Add(time.Hour * 24).Unix(),
	// })

	// 生成 JWT
	token, err := utils.GenerateToken(storedUser.ID)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	// tokenString, err := token.SignedString([]byte(secretKey))
	// if err != nil {
	// 	return "", errors.New("failed to generate token")
	// }
	return token, nil
}
