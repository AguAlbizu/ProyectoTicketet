package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleClient Role = "cliente"
	RoleAdmin  Role = "admin"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Email     string         `gorm:"type:varchar(150);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`
	Role      Role           `gorm:"type:enum('cliente','admin');default:'cliente';not null" json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
