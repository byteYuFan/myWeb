package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"myWeb/config"
	"time"
)

type Claims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成 token

func GenerateToken(id int64, username string) (string, error) {
	// 定义 token 的过期时间
	expireTime := time.Now().Add(config.ExpireDuration).Unix()

	// 创建一个自定义的 Claims
	claims := &Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "pogf",
		},
	}

	// 使用 JWT 签名算法生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 将 token 进行加盐加密
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return config.JwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}
}
