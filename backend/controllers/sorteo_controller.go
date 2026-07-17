package controllers

import (
	"net/http"
	"strconv"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

type SorteoController struct {
	sorteoService *services.SorteoService
}

func NewSorteoController(sorteoService *services.SorteoService) *SorteoController {
	return &SorteoController{sorteoService: sorteoService}
}

type createSorteoRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	ValorChance int    `json:"valor_chance" binding:"required"`
}

type buyChancesRequest struct {
	Cantidad int `json:"cantidad" binding:"required"`
}

// GetSorteoByEvent handles GET /api/events/:id/sorteo — público, sin token.
// Que el evento no tenga sorteo cargado es un estado normal, no un error: se responde
// 200 con body null en vez de 404, para no ensuciar la consola del front en cada evento
// sin sorteo (la mayoría). Un 404/500 acá queda reservado para fallos reales.
func (c *SorteoController) GetSorteoByEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	sorteo, err := c.sorteoService.GetSorteoByEventID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el sorteo"})
		return
	}
	ctx.JSON(http.StatusOK, sorteo)
}

// BuyChances handles POST /api/sorteos/:id/chances — requiere JWT.
func (c *SorteoController) BuyChances(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req buyChancesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cantidad es requerida"})
		return
	}

	chances, err := c.sorteoService.BuyChances(userID.(uint), uint(id), req.Cantidad)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "sorteo no encontrado" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, chances)
}

// GetMyChances handles GET /api/sorteos/:id/my-chances — requiere JWT.
func (c *SorteoController) GetMyChances(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	count, err := c.sorteoService.GetMyChancesCount(userID.(uint), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las chances"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"chances": count})
}

// CreateSorteo handles POST /api/admin/events/:id/sorteo — requiere JWT + rol admin.
func (c *SorteoController) CreateSorteo(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var req createSorteoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "nombre y valor_chance son requeridos"})
		return
	}

	sorteo, err := c.sorteoService.CreateSorteo(uint(id), req.Nombre, req.ValorChance)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "evento no encontrado" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, sorteo)
}

// ListSorteosAdmin handles GET /api/admin/sorteos — requiere JWT + rol admin.
func (c *SorteoController) ListSorteosAdmin(ctx *gin.Context) {
	sorteos, err := c.sorteoService.GetSorteosConEvento()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los sorteos"})
		return
	}
	ctx.JSON(http.StatusOK, sorteos)
}

// RunDraw handles POST /api/admin/sorteos/:id/draw — requiere JWT + rol admin.
func (c *SorteoController) RunDraw(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	winner, err := c.sorteoService.RunDraw(uint(id))
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "sorteo no encontrado" {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Sorteo realizado", "ganador": winner})
}

// GetSorteosByEvent handles GET /api/admin/events/:id/sorteos — requiere JWT + rol admin.
// Retorna el historial completo de sorteos del evento (activos y ya realizados).
func (c *SorteoController) GetSorteosByEvent(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	sorteos, err := c.sorteoService.GetSorteosByEventID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los sorteos del evento"})
		return
	}
	ctx.JSON(http.StatusOK, sorteos)
}

// GetChanceSummary handles GET /api/admin/sorteos/:id/chances — requiere JWT + rol admin.
// Retorna, para el sorteo, cada usuario participante con la cantidad de chances que compró.
func (c *SorteoController) GetChanceSummary(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	summary, err := c.sorteoService.GetChanceSummary(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los participantes del sorteo"})
		return
	}
	ctx.JSON(http.StatusOK, summary)
}
