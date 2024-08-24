package middlewares

import (
	"app/common/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {

	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Missing authorization token",
		})
		return
	}

	userId, err := jwt.Verify(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization token",
		})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
