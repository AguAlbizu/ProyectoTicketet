package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

// UserDAO handles all database operations for the User model.
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO creates a new UserDAO with the provided GORM instance.
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

// Create persists a new user record.
func (d *UserDAO) Create(user *domain.User) error {
	// TODO: d.db.Create(user)
	return nil
}

// FindByEmail returns the user matching the given email, or an error if not found.
func (d *UserDAO) FindByEmail(email string) (*domain.User, error) {
	// TODO: d.db.Where("email = ?", email).First(&user)
	return nil, nil
}

// FindByID returns the user matching the given ID.
func (d *UserDAO) FindByID(id uint) (*domain.User, error) {
	// TODO: d.db.First(&user, id)
	return nil, nil
}
