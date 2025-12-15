package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"metanode.com/homework/server/config"
)

// Claims 自定义 JWT 结构体
type Claims struct {
	UserID uint `json:"userId"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT
func GenerateToken(userID uint) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour) // 24小时过期
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwtSecret := []byte(config.GetJWTSecretKey())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析并验证 JWT
func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetJWTSecretKey()), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		return nil, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
