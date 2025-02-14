package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mingloo/web-api/internal/models"
	"mingloo/web-api/internal/repositories"
)

func GetUsers(c *gin.Context) {
	users := repositories.FetchUsers()
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repositories.SaveUser(user)
	c.JSON(http.StatusCreated, user)
}
