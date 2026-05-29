package services

import (
	"fmt"
	"ticketapp/domain"
	"ticketapp/utils"

	"gorm.io/gorm"
)

// AuthUserDAO define los métodos de persistencia requeridos por AuthService.
type AuthUserDAO interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
}

type AuthService struct {
	userDAO AuthUserDAO
}

func NewAuthService(userDAO AuthUserDAO) *AuthService {
	return &AuthService{userDAO: userDAO}
}

// Register hashea la contraseña y crea un nuevo usuario con rol "cliente".
// Retorna error si el email ya está registrado.
func (s *AuthService) Register(nombre, email, password string) error {
	_, err := s.userDAO.GetUserByEmail(email)
	if err == nil {
		return fmt.Errorf("el email ya está registrado")
	}
	if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error al verificar email: %w", err)
	}

	user := &domain.User{
		Nombre:   nombre,
		Email:    email,
		Password: utils.HashPassword(password),
		Rol:      "cliente",
	}
	return s.userDAO.CreateUser(user)
}

// Login verifica credenciales y retorna un JWT firmado.
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userDAO.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("credenciales inválidas")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", fmt.Errorf("credenciales inválidas")
	}

	token, err := utils.GenerateToken(user.ID, user.Rol, user.Email)
	if err != nil {
		return "", fmt.Errorf("error al generar token: %w", err)
	}
	return token, nil
}

// GetUserByEmail expone la búsqueda de usuario para el controller de login.
func (s *AuthService) GetUserByEmail(email string) (*domain.User, error) {
	return s.userDAO.GetUserByEmail(email)
}
