package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims holds the JWT payload.
// NOTE (entrega parcial): el token se usa SOLO para autenticación (verificar identidad),
// no para autorización (verificar permisos por rol). El campo Role está incluido
// para que la entrega final pueda extender esta struct sin cambios breaking.
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates and signs a JWT for the given user ID and role.
// Returns the signed token string or an error.
func GenerateToken(userID uint, role string) (string, error) {
	// TODO (entrega final): leer JWT_EXPIRATION_HOURS desde env para el tiempo de expiración
	// TODO: build Claims{UserID, Role, RegisteredClaims{ExpiresAt: now + expiration}}
	// TODO: jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return "", nil
}

// ValidateToken parses and validates the given JWT string.
// Returns the Claims if the token is valid and not expired.
// NOTE (entrega parcial): solo verifica firma y expiración, NO verifica el rol.
func ValidateToken(tokenString string) (*Claims, error) {
	// TODO: leer JWT_SECRET desde env
	// TODO: jwt.ParseWithClaims(tokenString, &Claims{}, keyFunc)
	// TODO: verificar que el método de firma sea HMAC
	// TODO: retornar claims si válido, error si expirado o firma inválida
	return nil, nil
}
