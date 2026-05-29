package controllers

import (
	"net/http"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

// TicketController exposes HTTP handlers for ticket-related endpoints.
// Todas las rutas requieren autenticación JWT.
type TicketController struct {
	ticketService *services.TicketService
}

// NewTicketController creates a new TicketController.
func NewTicketController(ticketService *services.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

// BuyTicket handles POST /api/tickets
// Compra una entrada para un evento en nombre del usuario autenticado.
func (c *TicketController) BuyTicket(ctx *gin.Context) {
	// TODO: obtener userID desde ctx.Get("userID") (seteado por el middleware JWT)
	// TODO: bindear JSON body a struct con campo event_id
	// TODO: llamar c.ticketService.Purchase(userID, eventID)
	// TODO: retornar 201 con el ticket creado, o 400/409 si no hay cupo
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// GetMyTickets handles GET /api/tickets/my-tickets
// Retorna todas las entradas del usuario autenticado.
func (c *TicketController) GetMyTickets(ctx *gin.Context) {
	// TODO: obtener userID desde ctx.Get("userID")
	// TODO: llamar c.ticketService.GetByUser(userID)
	// TODO: retornar 200 con la lista de tickets
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// CancelTicket handles DELETE /api/tickets/:id
// Cancela una entrada activa del usuario autenticado.
func (c *TicketController) CancelTicket(ctx *gin.Context) {
	// TODO: parsear ticketID desde ctx.Param("id")
	// TODO: obtener userID desde ctx.Get("userID")
	// TODO: llamar c.ticketService.Cancel(ticketID, userID)
	// TODO: retornar 204, o 403 si el ticket no pertenece al usuario, o 404 si no existe
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// TransferTicket handles PUT /api/tickets/:id/transfer
// Transfiere una entrada activa a otro usuario identificado por email.
func (c *TicketController) TransferTicket(ctx *gin.Context) {
	// TODO: parsear ticketID desde ctx.Param("id")
	// TODO: obtener ownerID desde ctx.Get("userID")
	// TODO: bindear JSON body a struct con campo target_email
	// TODO: llamar c.ticketService.Transfer(ticketID, ownerID, targetEmail)
	// TODO: retornar 200, o 403/404 según corresponda
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
