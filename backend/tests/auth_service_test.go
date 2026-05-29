package tests

// Objetivo de cobertura para la entrega parcial: >= 40% en servicios y controladores.
// Correr con: go test ./tests/... -cover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRegisterUser_Success
// Caso de éxito: registrar un usuario con datos válidos debe retornar el usuario creado sin error.
func TestRegisterUser_Success(t *testing.T) {
	// TODO: inicializar un UserDAO con DB en memoria (sqlite) o mock
	// TODO: crear AuthService con ese DAO
	// TODO: llamar service.Register con name, email y password válidos
	// TODO: assert que no haya error
	// TODO: assert que el usuario retornado tenga el email y rol "cliente"
	assert.True(t, true, "placeholder — implementar test body")
}

// TestRegisterUser_DuplicateEmail
// Caso de error: registrar con un email ya existente debe retornar error.
func TestRegisterUser_DuplicateEmail(t *testing.T) {
	// TODO: seedear un usuario con email "test@example.com"
	// TODO: llamar service.Register con el mismo email
	// TODO: assert que se retorne un error de duplicado
	assert.True(t, true, "placeholder — implementar test body")
}

// TestLogin_Success
// Caso de éxito: credenciales correctas deben retornar un token JWT no vacío.
func TestLogin_Success(t *testing.T) {
	// TODO: seedear usuario con email y password (hasheada)
	// TODO: llamar service.Login con las credenciales correctas
	// TODO: assert que no haya error
	// TODO: assert que el token retornado sea no vacío
	assert.True(t, true, "placeholder — implementar test body")
}

// TestLogin_WrongPassword
// Caso de error: password incorrecta debe retornar error de credenciales inválidas.
func TestLogin_WrongPassword(t *testing.T) {
	// TODO: seedear usuario
	// TODO: llamar service.Login con password incorrecta
	// TODO: assert que se retorne un error
	assert.True(t, true, "placeholder — implementar test body")
}
