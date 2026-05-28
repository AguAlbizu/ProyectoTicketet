package services

import (
	"ticketapp/dao"
	"ticketapp/domain"
	"ticketapp/utils"
)

// AuthService handles registration, login, and token management.
type AuthService struct {
	userDAO *dao.UserDAO
}

// NewAuthService creates a new AuthService with its required dependencies.
func NewAuthService(userDAO *dao.UserDAO) *AuthService {
	return &AuthService{userDAO: userDAO}
}

// RegisterInput holds the data required to register a new user.
type RegisterInput struct {
	Name     string
	Email    string
	Password string
	Role     domain.Role
}

// Register creates a new user account after validating the input.
// Returns the created user or an error (e.g. email already taken).
func (s *AuthService) Register(input RegisterInput) (*domain.User, error) {
	// TODO: validate that email is not already registered (userDAO.FindByEmail)
	// TODO: hash password with utils.HashPassword
	// TODO: create and persist user via userDAO.Create
	// TODO: return user
	_ = utils.HashPassword
	return nil, nil
}

// LoginOutput holds the data returned after a successful login.
type LoginOutput struct {
	Token string
	User  *domain.User
}

// Login authenticates a user and returns a signed JWT on success.
func (s *AuthService) Login(email, password string) (*LoginOutput, error) {
	// TODO: find user by email (userDAO.FindByEmail)
	// TODO: verify password with utils.CheckPassword
	// TODO: generate JWT with utils.GenerateToken(user.ID, string(user.Role))
	// TODO: return LoginOutput{Token, User}
	_ = utils.GenerateToken
	return nil, nil
}
