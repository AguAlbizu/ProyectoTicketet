package middleware

import (
	"net/http"
	"strings"
	"ticketapp/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware valida el JWT del header Authorization: Bearer <token>.
// Si es válido, inyecta userID, role y email en el contexto de Gin.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido. Usar: Bearer <token>"})
			return
		}

		claims, err := utils.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Set("email", claims.Email)
		c.Next()
	}
}

// AdminOnly corta la cadena con 403 si el rol del token (inyectado por AuthMiddleware) no es "admin".
// Debe usarse siempre después de AuthMiddleware en la cadena de handlers.
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Requiere permisos de administrador"})
			return
		}
		c.Next()
	}
}
