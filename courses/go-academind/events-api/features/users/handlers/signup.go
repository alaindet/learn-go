package handlers

import (
	"app/features/users/models"
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

	// event.UserID = 1 // TODO
	signedUpUser, err := user.Create()
	if err != nil {

		// TODO: Check for existing user

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
