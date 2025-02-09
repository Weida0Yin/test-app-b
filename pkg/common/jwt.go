package common

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"novel-app/pkg"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv(pkg.GetEnv("JWT_SECRET", "")))

// Claims 定义JWT载荷
type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT 生成 JWT
func GenerateJWT(userID int64) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(getJWTExpire()) * time.Second).Unix(), // 设置过期时间
			Issuer:    "novel-app",                                                        // 颁发者
		},
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并返回
	return token.SignedString(secretKey)
}

// ParseJWT 解析 JWT
func ParseJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 返回秘钥
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func getJWTExpire() int {
	expire := os.Getenv("JWT_EXP")
	if expire == "" {
		return 3600
	}
	return 3600
}
