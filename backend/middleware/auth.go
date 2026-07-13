package middleware

import (
	"net/http"
	"strings"
	"ticketapp/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware valida el JWT del header Authorization: Bearer <token>.
// Si es válido, inyecta userID, email y role en el contexto de Gin.
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
		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// RequireRole devuelve un middleware que verifica que el usuario autenticado
// tenga el rol especificado. Debe usarse después de AuthMiddleware.
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists || userRole.(string) != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Acceso denegado: permisos insuficientes"})
			return
		}
		c.Next()
	}
}
