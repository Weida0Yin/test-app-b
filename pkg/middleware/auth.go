package middleware

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"novel-app/internal/repo"
	"novel-app/pkg/common"
	"novel-app/pkg/response"
)

// JWTAuth JWT 中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, "Authorization header required")
			//c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "msg": "Authorization header required"})
			c.Abort()
			return
		}

		// 提取 JWT
		token := authHeader[len("Bearer "):]

		// 解析 JWT
		claims, err := common.ParseJWT(token)
		if err != nil {
			response.Fail(c, "invalid or expired token")
			//c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "msg": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 检查 Token 是否在 Redis 中
		redisKey := fmt.Sprintf("user_auth_%d", claims.UserID)
		cachedToken, err := repo.RdsClt.Get(context.Background(), redisKey).Result()
		if err != nil || cachedToken != token {
			response.Fail(c, "Token expired or invalid")
			c.Abort()
			return
		}

		// 将 user_id 放到上下文中，供后续使用
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
