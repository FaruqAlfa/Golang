package main

/*
AUTHENTICATION & AUTHORIZATION
Implementing Role-Based Access Control (RBAC) in Gin

Instructions: Create an HTTP server using the Gin framework that implements Role-Based Access Control (RBAC) to manage user access to different routes based on their role.

Instructions

Define User Roles: Users are assigned roles of admin, editor, or viewer.
admin can access /admin, /edit, and /view.
editor can access /edit and /view.
viewer can only access /view.
Implement Middleware: Write middleware in Gin that checks the user's role and permits access to routes accordingly.
Create Endpoints: Set up endpoints for /admin, /edit, and /view */


import (
	"net/http"
	"github.com/gin-gonic/gin"
)

//mendefinisikan peran admin, editor, dan viewer
const(
	RoleAdmin = "admin"
	RoleEditor = "editor"
	RoleViewer = "viewer"
)

//membuat struct untuk menyimpan data user
type User struct {
	Username string 
	Role	 string
}

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exist := c.Get("user")//mendapatkan user dari request
		if !exist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		currentUser := user.(User)//memeriksa peran user
		for _, role := range allowedRoles {
			if currentUser.Role == role {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
	}
}

func main() {
	router := gin.Default()
	simulaedUser := User{Username: "faruq", Role: RoleAdmin}

	router.Use(func(ctx *gin.Context) {
		ctx.Set("user", simulaedUser)
		ctx.Next()
	})

	adminGroup := router.Group("/admin")
	adminGroup.Use(RoleMiddleware(RoleAdmin))
	{
		adminGroup.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, admin!"})
		})
	}

	editGroup := router.Group("/edit")
	editGroup.Use(RoleMiddleware(RoleAdmin, RoleEditor))
	{
		editGroup.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, editor!"})
		})
	}

	viewGroup := router.Group("/view")
	viewGroup.Use(RoleMiddleware(RoleAdmin, RoleEditor, RoleViewer))
	{
		viewGroup.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello, viewer!"})
		})
	}

	router.Run(":8080")
}