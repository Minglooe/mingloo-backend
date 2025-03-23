package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEvents handles fetching events
// @Summary Retorna uma lista dos eventos
// @Description Retornar uma lista contendo todos os eventos ativos na aplicação
// @Tags Eventos
// @Produce json
// @Success 200 {object} object{success=boolean,data=[]models.Event}
// @Router /events [get]
func GetEvents(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "This route is not implemented yet.",
	})
}

// CreateEvent handles creating a new event
// @Summary Cadastra um novo evento
// @Description Cadastra um novo evento na aplicação com o status ativo
// @Tags Eventos
// @Accept json
// @Produce json
// @Success 200 {object} object{success=boolean,message=string}
// @Router /events [post]
func CreateEvent(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "This route is not implemented yet.",
	})
}
