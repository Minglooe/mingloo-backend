package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetEvents handles fetching events
// @Summary Retorna uma lista dos eventos
// @Description Retornar uma lista contendo todos os eventos ativos na aplicação.
// @Tags Eventos
// @Produce json
// @Success 200 {object} http_responses.ResponseWrapper{data=[]models.Event}
// @Router /events [get]
func GetEvents(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "This route is not implemented yet.",
	})
}

// CreateEvent handles creating a new event
// @Summary Cadastra um novo evento
// @Description Cadastra um novo evento na aplicação com o status ativo.
// @Tags Eventos
// @Accept json
// @Produce json
// @Success 200 {object} http_responses.ResponseWrapper{data=string}
// @Router /events [post]
func CreateEvent(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "This route is not implemented yet.",
	})
}

// @Summary Atualiza um evento
// @Description Atualize um evento na aplicação por meio do seu ID.
// @Tags Eventos
// @Accept json
// @Produce json
// @Param id path int true "Id do Evento a ser atualizado"
// @Success 200 {object} http_responses.ResponseWrapper{data=[]models.Event}
// @Router /events/{id} [put]
func UpdateEvent(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "This route is not implemented yet.",
	})
}

// @Summary Deleta um evento
// @Description Atualize um evento da aplicação por meio do seu ID.
// @Tags Eventos
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} http_responses.ResponseWrapper{data=string}
// @Router /events/{id} [delete]
func DeleteEvent(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"message": "This route is not implemented yet.",
	})
}
