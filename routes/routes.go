package routes

import (
	"github.com/gin-gonic/gin"
	"redis-like/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	// Add more routes as needed
}
