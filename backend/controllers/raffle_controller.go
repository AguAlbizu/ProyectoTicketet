package controllers

import (
	"net/http"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

// RaffleController exposes HTTP handlers for raffle-related endpoints.
type RaffleController struct {
	raffleService *services.RaffleService
}

// NewRaffleController creates a new RaffleController.
func NewRaffleController(raffleService *services.RaffleService) *RaffleController {
	return &RaffleController{raffleService: raffleService}
}

// RegisterRoutes wires the raffle endpoints onto the given router group.
// GET  /api/raffles/event/:eventId      - public: get raffle for an event
// POST /api/raffles/:id/chances         - auth: buy chances
// POST /api/raffles/:id/draw            - admin only: execute draw
// POST /api/raffles                     - admin only: create raffle for event
func (c *RaffleController) RegisterRoutes(rg *gin.RouterGroup, authMiddleware gin.HandlerFunc, adminMiddleware gin.HandlerFunc) {
	raffles := rg.Group("/raffles")
	raffles.GET("/event/:eventId", c.GetByEvent)

	authRaffles := raffles.Group("")
	authRaffles.Use(authMiddleware)
	authRaffles.POST("/:id/chances", c.BuyChances)

	adminRaffles := raffles.Group("")
	adminRaffles.Use(authMiddleware, adminMiddleware)
	adminRaffles.POST("", c.Create)
	adminRaffles.POST("/:id/draw", c.Draw)
}

// GetByEvent handles GET /api/raffles/event/:eventId.
func (c *RaffleController) GetByEvent(ctx *gin.Context) {
	// TODO: parse eventID from ctx.Param("eventId")
	// TODO: call raffleService.GetByEvent(eventID)
	// TODO: return 200 with raffle JSON or 404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// BuyChances handles POST /api/raffles/:id/chances.
func (c *RaffleController) BuyChances(ctx *gin.Context) {
	// TODO: parse raffleID, get userID from context
	// TODO: bind JSON body to {quantity}
	// TODO: call raffleService.BuyChances(userID, raffleID, quantity)
	// TODO: return 200 with entry JSON or 400/404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Create handles POST /api/raffles (admin only).
func (c *RaffleController) Create(ctx *gin.Context) {
	// TODO: bind JSON body to RaffleInput
	// TODO: call raffleService.Create(input)
	// TODO: return 201 with raffle JSON or 400
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Draw handles POST /api/raffles/:id/draw (admin only).
func (c *RaffleController) Draw(ctx *gin.Context) {
	// TODO: parse raffleID from ctx.Param("id")
	// TODO: call raffleService.Draw(raffleID)
	// TODO: return 200 with winner user JSON or 400/404
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
