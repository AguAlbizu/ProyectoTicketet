package controllers

import (
	"net/http"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

// TicketController exposes HTTP handlers for ticket-related endpoints.
type TicketController struct {
	ticketService *services.TicketService
}

// NewTicketController creates a new TicketController.
func NewTicketController(ticketService *services.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

// RegisterRoutes wires the ticket endpoints onto the given router group.
// All routes require authentication (client role).
// POST   /api/tickets              - purchase a ticket for an event
// GET    /api/tickets/my           - list my tickets
// DELETE /api/tickets/:id          - cancel a ticket
// POST   /api/tickets/:id/transfer - transfer ticket to another user
func (c *TicketController) RegisterRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	tickets := rg.Group("/tickets")
	tickets.Use(authMiddleware)
	tickets.POST("", c.Purchase)
	tickets.GET("/my", c.GetMyTickets)
	tickets.DELETE("/:id", c.Cancel)
	tickets.POST("/:id/transfer", c.Transfer)
}

// Purchase handles POST /api/tickets.
func (c *TicketController) Purchase(ctx *gin.Context) {
	// TODO: get authenticated userID from context (set by JWT middleware)
	// TODO: bind JSON body to {event_id}
	// TODO: call ticketService.Purchase(userID, eventID)
	// TODO: return 201 with ticket JSON or 400/409 on error
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// GetMyTickets handles GET /api/tickets/my.
func (c *TicketController) GetMyTickets(ctx *gin.Context) {
	// TODO: get authenticated userID from context
	// TODO: call ticketService.GetByUser(userID)
	// TODO: return 200 with tickets JSON
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Cancel handles DELETE /api/tickets/:id.
func (c *TicketController) Cancel(ctx *gin.Context) {
	// TODO: parse ticketID, get userID from context
	// TODO: call ticketService.Cancel(ticketID, userID)
	// TODO: return 204 or 403/404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Transfer handles POST /api/tickets/:id/transfer.
func (c *TicketController) Transfer(ctx *gin.Context) {
	// TODO: parse ticketID, get ownerID from context
	// TODO: bind JSON body to {target_email}
	// TODO: call ticketService.Transfer(ticketID, ownerID, targetEmail)
	// TODO: return 200 or 403/404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
