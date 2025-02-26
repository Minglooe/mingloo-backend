package handlers

import (
	"net/http"

	"mingloo/web-api/internal/models"
	"mingloo/web-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary GET usuários
// @Schemes user users
// @Description retorna os usuários da aplicação
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /api/v1/helloworld [get]
func GetUsers(c *gin.Context) {
	users := repositories.FetchUsers()
	c.JSON(http.StatusOK, users)
}

// PingExample godoc
// @Summary GET usuários
// @Schemes
// @Description retorna os usuários da aplicação
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Sucesso
// @Router /api/v1/helloworld2 [get]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repositories.SaveUser(user)
	c.JSON(http.StatusCreated, user)
}
