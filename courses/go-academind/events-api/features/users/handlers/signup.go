package handlers

import (
	"app/features/users/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpUser(ctx *gin.Context) {

	var user models.UserModel
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot sign up user due to invalid data",
		})
		return
	}

	signedUpUser, err := user.Create()
	if err != nil {

		if errors.Is(err, models.ErrUserExists) {
			ctx.JSON(http.StatusConflict, gin.H{
				"message": fmt.Sprintf("User %q already exists", user.Email),
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Cannot create user on the database",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("User %q signed up", user.Email),
		"data":    signedUpUser,
	})
}
