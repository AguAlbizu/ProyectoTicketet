package controllers

import (
	"net/http"
	"strconv"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
	ticketService *services.TicketService
}

func NewTicketController(ticketService *services.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

type buyTicketRequest struct {
	EventID uint `json:"event_id" binding:"required"`
}

type transferRequest struct {
	TargetEmail string `json:"target_email" binding:"required"`
}

// BuyTicket handles POST /api/tickets
func (c *TicketController) BuyTicket(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var req buyTicketRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "event_id es requerido"})
		return
	}

	ticket, err := c.ticketService.BuyTicket(userID.(uint), req.EventID)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "evento no encontrado" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, ticket)
}

// GetMyTickets handles GET /api/tickets/my-tickets
func (c *TicketController) GetMyTickets(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	tickets, err := c.ticketService.GetMyTickets(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener entradas"})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

// CancelTicket handles DELETE /api/tickets/:id
func (c *TicketController) CancelTicket(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = c.ticketService.CancelTicket(uint(id), userID.(uint))
	if err != nil {
		switch err.Error() {
		case "entrada no encontrada":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "no tenés permiso para cancelar esta entrada":
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Entrada cancelada"})
}

// TransferTicket handles PUT /api/tickets/:id/transfer
func (c *TicketController) TransferTicket(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "target_email es requerido"})
		return
	}

	err = c.ticketService.TransferTicket(uint(id), userID.(uint), req.TargetEmail)
	if err != nil {
		switch err.Error() {
		case "usuario destino no encontrado", "entrada no encontrada":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "no tenés permiso para transferir esta entrada":
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Entrada transferida exitosamente"})
}
