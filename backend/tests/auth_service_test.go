package tests

// Objetivo de cobertura para la entrega parcial: >= 40% en servicios y controladores.
// Correr con: go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils -v

import (
	"os"
	"testing"
	"ticketapp/domain"
	"ticketapp/services"
	"ticketapp/utils"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// mockAuthUserDAO implementa services.AuthUserDAO para tests sin base de datos.
type mockAuthUserDAO struct {
	user      *domain.User
	createErr error
	findErr   error
}

func (m *mockAuthUserDAO) CreateUser(user *domain.User) error { return m.createErr }
func (m *mockAuthUserDAO) GetUserByEmail(email string) (*domain.User, error) {
	if m.findErr != nil {
		return nil, m.findErr
	}
	return m.user, nil
}
func (m *mockAuthUserDAO) GetUserByID(id uint) (*domain.User, error) { return m.user, nil }

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

// TestRegister_Success verifica que un usuario nuevo se registra correctamente.
func TestRegister_Success(t *testing.T) {
	mockDAO := &mockAuthUserDAO{findErr: gorm.ErrRecordNotFound}
	svc := services.NewAuthService(mockDAO)

	err := svc.Register("Juan", "juan@test.com", "password123")
	assert.NoError(t, err)
}

// TestRegister_DuplicateEmail verifica que registrar un email ya existente retorna error.
func TestRegister_DuplicateEmail(t *testing.T) {
	mockDAO := &mockAuthUserDAO{user: &domain.User{Email: "juan@test.com"}}
	svc := services.NewAuthService(mockDAO)

	err := svc.Register("Juan", "juan@test.com", "password123")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "el email ya está registrado")
}

// TestLogin_Success verifica que credenciales válidas retornan un JWT no vacío.
func TestLogin_Success(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-para-tests")
	defer os.Unsetenv("JWT_SECRET")

	hashed := utils.HashPassword("password123")
	mockDAO := &mockAuthUserDAO{
		user: &domain.User{IDUsers: 1, Email: "juan@test.com", Password: hashed, Rol: "cliente"},
	}
	svc := services.NewAuthService(mockDAO)

	token, err := svc.Login("juan@test.com", "password123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

// TestLogin_WrongPassword verifica que una contraseña incorrecta retorna error de credenciales.
func TestLogin_WrongPassword(t *testing.T) {
	hashed := utils.HashPassword("correcta")
	mockDAO := &mockAuthUserDAO{
		user: &domain.User{IDUsers: 1, Email: "juan@test.com", Password: hashed, Rol: "cliente"},
	}
	svc := services.NewAuthService(mockDAO)

	_, err := svc.Login("juan@test.com", "incorrecta")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "credenciales inválidas")
}

// TestLogin_UserNotFound verifica que un email inexistente retorna error de credenciales.
func TestLogin_UserNotFound(t *testing.T) {
	mockDAO := &mockAuthUserDAO{findErr: gorm.ErrRecordNotFound}
	svc := services.NewAuthService(mockDAO)

	_, err := svc.Login("noexiste@test.com", "cualquiera")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "credenciales inválidas")
}
