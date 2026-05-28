package utils

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims holds the JWT payload fields.
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken creates and signs a JWT for the given user ID and role.
// Returns the signed token string or an error.
func GenerateToken(userID uint, role string) (string, error) {
	// TODO: read JWT_SECRET and JWT_EXPIRATION_HOURS from env
	// TODO: build Claims{UserID, Role, RegisteredClaims{ExpiresAt}}
	// TODO: sign with jwt.SigningMethodHS256 and return tokenString
	return "", nil
}

// ValidateToken parses and validates the given JWT string.
// Returns the Claims if valid, or an error if expired/invalid.
func ValidateToken(tokenString string) (*Claims, error) {
	// TODO: parse tokenString using jwt.ParseWithClaims
	// TODO: verify signing method is HMAC
	// TODO: return claims if valid, else return error
	return nil, nil
}
