package main

/*
Exercise 3: Authentication Middleware

Instruction: Implement an authentication middleware that checks for an Authorization header.
If the header matches Bearer valid-token, allow access; otherwise, deny with 401 Unauthorized.
*/

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Autohorozation")

	if authHeader != "Bearer valid-token"{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}

	c.Next()
}

func main() {
	router := gin.Default()

	router.Use(AuthMiddleware)

	router.GET("/protected", func (c *gin.Context) {
		c.String(http.StatusOK, "You are authorized")
	})

	fmt.Println("Listening on port 8080")
	router.Run(":8080")
}