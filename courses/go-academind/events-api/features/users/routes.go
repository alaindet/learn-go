package users

import (
	"app/features/users/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {
	// TODO: Grouping
	server.POST("/signup", handlers.SignUpUser)
	server.POST("/signin", handlers.SignInUser)
}
