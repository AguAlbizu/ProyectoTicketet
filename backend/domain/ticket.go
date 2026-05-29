package domain

import (
	"time"
)

type Ticket struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint      `gorm:"not null;index" json:"user_id"`
	EventID      uint      `gorm:"not null;index" json:"event_id"`
	Status       string    `gorm:"type:varchar(20);default:'activo';not null" json:"estado"`
	PurchaseDate time.Time `gorm:"not null" json:"fecha_compra"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	// Associations
	User  User  `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Event Event `gorm:"foreignKey:EventID" json:"event,omitempty"`
}
