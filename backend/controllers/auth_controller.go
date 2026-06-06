package controllers

import (
	"net/http"
	"strings"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

type registerRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Register handles POST /api/auth/register
func (c *AuthController) Register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Faltan campos requeridos: nombre, email, password"})
		return
	}

	err := c.authService.Register(req.Nombre, req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "ya está registrado") {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar usuario"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario registrado exitosamente"})
}

// Login handles POST /api/auth/login
func (c *AuthController) Login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email y password son requeridos"})
		return
	}

	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	user, _ := c.authService.GetUserByEmail(req.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id_users": user.IDUsers,
			"nombre":   user.Nombre,
			"email":    user.Email,
			"rol":      user.Rol,
		},
	})
}
