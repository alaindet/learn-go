package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func requestBodyExample(c *gin.Context) {
	fnameQuery := c.Query("fname")
	fnameBody := c.PostForm("fname")
	lnameBody := c.PostForm("lname")

	fmt.Printf("fnameQuery: (%T) %v\n", fnameQuery, fnameQuery)
	fmt.Printf("fnameBody: (%T) %v\n", fnameBody, fnameBody)
	fmt.Printf("lnameBody: (%T) %v\n", lnameBody, lnameBody)

	c.String(200, "Body test endpoint (check logs)")
}
