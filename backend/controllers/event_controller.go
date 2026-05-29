package controllers

import (
	"net/http"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

// EventController exposes HTTP handlers for event-related endpoints.
type EventController struct {
	eventService *services.EventService
}

// NewEventController creates a new EventController.
func NewEventController(eventService *services.EventService) *EventController {
	return &EventController{eventService: eventService}
}

// GetEvents handles GET /api/events
// Acepta query param opcional: ?categoria=<valor>
// Retorna la lista de eventos activos, filtrada por categoría si se indica.
func (c *EventController) GetEvents(ctx *gin.Context) {
	// TODO: leer query param "categoria" con ctx.Query("categoria")
	// TODO: llamar c.eventService.GetAll(categoria)
	// TODO: retornar 200 con la lista de eventos en JSON
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// GetEventByID handles GET /api/events/:id
// Retorna el detalle de un evento por su ID.
func (c *EventController) GetEventByID(ctx *gin.Context) {
	// TODO: parsear id uint desde ctx.Param("id")
	// TODO: llamar c.eventService.GetByID(id)
	// TODO: retornar 200 con el evento o 404 si no existe
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// TODO (entrega final): agregar handlers para crear, editar y cancelar eventos (admin)
