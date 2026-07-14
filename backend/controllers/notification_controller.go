package controllers

import (
	"net/http"
	"strconv"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationService *services.NotificationService
}

func NewNotificationController(notificationService *services.NotificationService) *NotificationController {
	return &NotificationController{notificationService: notificationService}
}

// GetMyNotifications handles GET /api/notifications — requiere JWT.
func (c *NotificationController) GetMyNotifications(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	notifications, err := c.notificationService.GetMyNotifications(userID.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las notificaciones"})
		return
	}
	ctx.JSON(http.StatusOK, notifications)
}

// MarkAsRead handles PUT /api/notifications/:id/read — requiere JWT.
func (c *NotificationController) MarkAsRead(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.notificationService.MarkAsRead(uint(id), userID.(uint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al marcar la notificación como leída"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Notificación marcada como leída"})
}

// MarkAllAsRead handles PUT /api/notifications/read-all — requiere JWT.
func (c *NotificationController) MarkAllAsRead(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	if err := c.notificationService.MarkAllAsRead(userID.(uint)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al marcar las notificaciones como leídas"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Notificaciones marcadas como leídas"})
}
