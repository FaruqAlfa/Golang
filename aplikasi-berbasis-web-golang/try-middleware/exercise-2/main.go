package main


/*Exercise 2: Route-Specific Middleware

Instruction: Add a route-specific middleware to handle a new route /admin. 
This middleware should check for a query parameter admin=true. If not present, respond with status code 403 Forbidden.

*/
import (
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.Query("admin") != "true" {
            c.AbortWithStatus(403) // Forbidden access
            return
        }
        c.Next()
    }
}

func main() {
    router := gin.Default()
    router.Use(AdminOnly())
    router.GET("/admin", func(c *gin.Context) {
        c.String(200, "Welcome Admin")
    })
    router.Run()
}