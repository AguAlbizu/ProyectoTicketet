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

// Register handles POST /api/auth/register
// Registra un nuevo usuario con rol "cliente" por defecto.
func (c *AuthController) Register(ctx *gin.Context) {
	// TODO: bindear JSON body a struct con campos name, email, password
	// TODO: llamar c.authService.Register(input)
	// TODO: retornar 201 con el usuario creado, o 400 si faltan campos, o 409 si el email ya existe
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}

// Login handles POST /api/auth/login
// Autentica un usuario existente y retorna un JWT.
func (c *AuthController) Login(ctx *gin.Context) {
	// TODO: bindear JSON body a struct con campos email, password
	// TODO: llamar c.authService.Login(email, password)
	// TODO: retornar 200 con { token, user }, o 401 si las credenciales son inválidas
	ctx.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented"})
}
