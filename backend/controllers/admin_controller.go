package controllers

import (
	"net/http"
	"strconv"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *services.AdminEventService
	authService  *services.AuthService
}

func NewAdminController(adminService *services.AdminEventService, authService *services.AuthService) *AdminController {
	return &AdminController{adminService: adminService, authService: authService}
}

type createAdminRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PromoteToAdmin handles PUT /api/admin/users/promote
func (c *AdminController) PromoteToAdmin(ctx *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El campo email es requerido"})
		return
	}

	if err := c.authService.PromoteToAdmin(req.Email); err != nil {
		switch err.Error() {
		case "usuario no encontrado":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "el usuario ya es administrador":
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el rol"})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario promovido a administrador exitosamente"})
}

// CreateAdmin handles POST /api/admin/users
func (c *AdminController) CreateAdmin(ctx *gin.Context) {
	var req createAdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faltan campos requeridos: nombre, email, password"})
		return
	}

	if err := c.authService.RegisterAdmin(req.Nombre, req.Email, req.Password); err != nil {
		if err.Error() == "el email ya está registrado" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear administrador"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Administrador creado exitosamente"})
}

// GetAllEvents handles GET /api/admin/events
func (c *AdminController) GetAllEvents(ctx *gin.Context) {
	events, err := c.adminService.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener eventos"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

// CreateEvent handles POST /api/admin/events
func (c *AdminController) CreateEvent(ctx *gin.Context) {
	var input services.CreateEventInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := c.adminService.CreateEvent(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

// UpdateEvent handles PUT /api/admin/events/:id
func (c *AdminController) UpdateEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var input services.UpdateEventInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := c.adminService.UpdateEvent(uint(id), input)
	if err != nil {
		switch err.Error() {
		case "evento no encontrado":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "no se puede modificar un evento cancelado":
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, event)
}

// CancelEvent handles DELETE /api/admin/events/:id
func (c *AdminController) CancelEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.adminService.CancelEvent(uint(id)); err != nil {
		switch err.Error() {
		case "evento no encontrado", "el evento ya está cancelado":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Evento cancelado exitosamente"})
}

// GetEventReport handles GET /api/admin/events/:id/report
func (c *AdminController) GetEventReport(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	report, err := c.adminService.GetEventReport(uint(id))
	if err != nil {
		switch err.Error() {
		case "evento no encontrado":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, report)
}
