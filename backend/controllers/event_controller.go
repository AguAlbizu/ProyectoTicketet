package controllers

import (
	"net/http"
	"strconv"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService *services.EventService
}

func NewEventController(eventService *services.EventService) *EventController {
	return &EventController{eventService: eventService}
}

// GetEvents handles GET /api/events?categoria=xxx
func (c *EventController) GetEvents(ctx *gin.Context) {
	categoria := ctx.Query("categoria")
	events, err := c.eventService.GetEvents(categoria)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener eventos"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// GetEventByID handles GET /api/events/:id
func (c *EventController) GetEventByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	event, err := c.eventService.GetEventByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

// TODO (entrega final): agregar handlers para crear, editar y cancelar eventos (admin)
