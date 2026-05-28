package controllers

import (
	"net/http"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

// AuthController exposes HTTP handlers for authentication endpoints.
type AuthController struct {
	authService *services.AuthService
}

// NewAuthController creates a new AuthController.
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// RegisterRoutes wires the auth endpoints onto the given router group.
// POST /api/auth/register
// POST /api/auth/login
func (c *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	auth.POST("/register", c.Register)
	auth.POST("/login", c.Login)
}

// Register handles POST /api/auth/register.
func (c *AuthController) Register(ctx *gin.Context) {
	// TODO: bind JSON body to RegisterInput struct
	// TODO: call authService.Register(input)
	// TODO: return 201 with user JSON or 400/409 on error
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Login handles POST /api/auth/login.
func (c *AuthController) Login(ctx *gin.Context) {
	// TODO: bind JSON body to {email, password}
	// TODO: call authService.Login(email, password)
	// TODO: return 200 with {token, user} or 401 on invalid credentials
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
