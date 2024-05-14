package main

/*
Exercise 5: Content-Type Enforcement Middleware

Instruction: Write middleware that ensures only JSON content is accepted for POST requests on /json-only.
If the content-type is not application/json, return an error.

*/

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ensureJSONContentType(c *gin.Context) {
	//check header content type
	contentType := c.Request.Header.Get("Content-Type")

	if contentType != "application/json" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid content type, Must be application/json"})
		return
	}
}

func main() {
	router := gin.Default()

	//add middleware for JSON
	router.POST("/json-only", ensureJSONContentType, func(c *gin.Context) {
		c.String(http.StatusOK, "received JSON content")
	})

	fmt.Println("Listening on port 8080")
	router.Run(":8080")
}