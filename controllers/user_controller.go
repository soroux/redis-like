package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	// Logic to get users
	c.JSON(http.StatusOK, gin.H{"message": "Get Users"})
}
