package main


/*
Exercise 1: Basic Logging Middleware

Instruction: Create a basic Gin server with a logging middleware that logs the request method and 
path for every incoming request. Implement a simple GET route / that returns "Hello, World!".
*/
import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
)

func main() {
	router := gin.Default()

	router.Use(func (c *gin.Context){
		log.Printf("%s %s", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println("Listening on port 8080")
	router.Run(":8080")
}