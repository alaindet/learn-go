package users

import (
	"app/features/users/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(routes *gin.RouterGroup) {
	routes.POST("/signup", handlers.SignUpUser)
	routes.POST("/signin", handlers.SignInUser)
}
