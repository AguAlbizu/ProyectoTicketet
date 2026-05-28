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

// RegisterRoutes wires the event endpoints onto the given router group.
// GET    /api/events          - public
// GET    /api/events/:id      - public
// POST   /api/events          - admin only
// PUT    /api/events/:id      - admin only
// DELETE /api/events/:id      - admin only
func (c *EventController) RegisterRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc, adminMiddleware gin.HandlerFunc) {
	events := rg.Group("/events")
	events.GET("", c.GetAll)
	events.GET("/:id", c.GetByID)

	adminEvents := events.Group("")
	adminEvents.Use(authMiddleware, adminMiddleware)
	adminEvents.POST("", c.Create)
	adminEvents.PUT("/:id", c.Update)
	adminEvents.DELETE("/:id", c.Cancel)
}

// GetAll handles GET /api/events with optional ?category= filter.
func (c *EventController) GetAll(ctx *gin.Context) {
	// TODO: read optional query param "category"
	// TODO: call eventService.GetAll(category)
	// TODO: return 200 with events JSON
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// GetByID handles GET /api/events/:id.
func (c *EventController) GetByID(ctx *gin.Context) {
	// TODO: parse uint id from ctx.Param("id")
	// TODO: call eventService.GetByID(id)
	// TODO: return 200 or 404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Create handles POST /api/events (admin only).
func (c *EventController) Create(ctx *gin.Context) {
	// TODO: bind JSON body to EventInput
	// TODO: call eventService.Create(input)
	// TODO: return 201 with event JSON or 400 on validation error
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Update handles PUT /api/events/:id (admin only).
func (c *EventController) Update(ctx *gin.Context) {
	// TODO: parse id, bind JSON body to EventInput
	// TODO: call eventService.Update(id, input)
	// TODO: return 200 or 404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Cancel handles DELETE /api/events/:id (admin only).
func (c *EventController) Cancel(ctx *gin.Context) {
	// TODO: parse id from ctx.Param("id")
	// TODO: call eventService.Cancel(id)
	// TODO: return 204 or 404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
