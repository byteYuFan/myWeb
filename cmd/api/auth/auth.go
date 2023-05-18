package auth

import (
	"github.com/gin-gonic/gin"
	"log"
	"myWeb/pkg/jwt"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": 100700,
				"description": "token错误",
			})
			ctx.Abort()
			return
		}
		user, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"status_code": 100701,
				"description": "token不可达",
			})
			ctx.Abort()
			return
		}
		log.Println("--------------", user.ID, user.Username)
		ctx.Set("id", user.ID)
		ctx.Set("username", user.Username)
		ctx.Next()
	}
}
