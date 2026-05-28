package domain

import (
	"time"

	"gorm.io/gorm"
)

type EventStatus string

const (
	EventStatusActive    EventStatus = "activo"
	EventStatusCancelled EventStatus = "cancelado"
)

type Event struct {
	ID               uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Title            string         `gorm:"type:varchar(200);not null" json:"title"`
	Description      string         `gorm:"type:text" json:"description"`
	Date             time.Time      `gorm:"not null" json:"date"`
	Duration         int            `gorm:"not null" json:"duration_minutes"`
	Capacity         int            `gorm:"not null" json:"capacity"`
	AvailableTickets int            `gorm:"not null" json:"available_tickets"`
	Category         string         `gorm:"type:varchar(100)" json:"category"`
	ImageURL         string         `gorm:"type:varchar(500)" json:"image_url"`
	Status           EventStatus    `gorm:"type:enum('activo','cancelado');default:'activo';not null" json:"status"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
