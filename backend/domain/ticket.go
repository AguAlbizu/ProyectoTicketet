package domain

import (
	"time"

	"gorm.io/gorm"
)

type TicketStatus string

const (
	TicketStatusActive      TicketStatus = "activo"
	TicketStatusCancelled   TicketStatus = "cancelado"
	TicketStatusTransferred TicketStatus = "transferido"
)

type Ticket struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint           `gorm:"not null;index" json:"user_id"`
	EventID      uint           `gorm:"not null;index" json:"event_id"`
	Status       TicketStatus   `gorm:"type:enum('activo','cancelado','transferido');default:'activo';not null" json:"status"`
	PurchaseDate time.Time      `gorm:"not null" json:"purchase_date"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Event Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
}
