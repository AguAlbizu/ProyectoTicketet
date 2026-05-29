package tests

// Objetivo de cobertura para la entrega parcial: >= 40% en servicios y controladores.
// Correr con: go test ./tests/... -v -cover

import (
	"os"
	"testing"
	"ticketapp/utils"

	"github.com/stretchr/testify/assert"
)

// TestHashPassword_ReturnsHash verifica que HashPassword retorna un hash no vacío de 64 chars (SHA-256 hex).
func TestHashPassword_ReturnsHash(t *testing.T) {
	hash := utils.HashPassword("password123")
	assert.NotEmpty(t, hash)
	assert.Len(t, hash, 64, "SHA-256 en hex debe tener 64 caracteres")
}

// TestCheckPassword_CorrectPassword verifica que la contraseña correcta pasa la validación.
func TestCheckPassword_CorrectPassword(t *testing.T) {
	hash := utils.HashPassword("miPassword")
	assert.True(t, utils.CheckPassword("miPassword", hash))
}

// TestCheckPassword_WrongPassword verifica que una contraseña incorrecta falla la validación.
func TestCheckPassword_WrongPassword(t *testing.T) {
	hash := utils.HashPassword("miPassword")
	assert.False(t, utils.CheckPassword("otraPassword", hash))
}

// TestGenerateAndValidateToken verifica que se puede generar un token y luego validarlo correctamente.
func TestGenerateAndValidateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test_secret_para_tests")
	os.Setenv("JWT_EXPIRATION_HOURS", "24")

	token, err := utils.GenerateToken(42, "cliente", "test@example.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims, err := utils.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, uint(42), claims.UserID)
	assert.Equal(t, "cliente", claims.Role)
	assert.Equal(t, "test@example.com", claims.Email)
}
