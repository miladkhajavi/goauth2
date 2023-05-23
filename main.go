package main

import (
	"goauth/controller"
	"goauth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	// Create a group, or base url for all routes

	// Welcome
	// r.GET("/", func(c *gin.Context) {
	// 	// c.String(http.StatusOK, "welcome")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"response":"welcome",
	// 	})
	// })
	r.GET("/google/login", controller.GoogleLogin)
	r.GET("/google/callback", controller.GoogleCallback)
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
