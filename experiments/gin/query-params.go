package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func queryParamsExample(c *gin.Context) {
	defaultFName := "This MUST exist"
	fname := c.DefaultQuery("firstname", defaultFName)
	lname := c.Query("lastname")

	if fname == defaultFName {
		c.String(400, "Please provide firstname and lastname query parameters")
		return
	}

	message := fmt.Sprintf("Hello, %s %s", fname, lname)
	c.String(200, message)
}
