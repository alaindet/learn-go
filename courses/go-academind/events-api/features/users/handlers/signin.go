package handlers

import (
	"app/features/users/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInUser(ctx *gin.Context) {

	// Parse send credentials
	var user models.UserModel
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot sign in user due to invalid data",
		})
		return
	}

	jwt, err := user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	// TODO: Create JWT

	ctx.JSON(http.StatusCreated, gin.H{
		"message":      fmt.Sprintf("Signed in with %s", user.Email),
		"access_token": jwt,
	})
}
