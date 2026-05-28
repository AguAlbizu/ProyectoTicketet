package domain

import (
	"time"

	"gorm.io/gorm"
)

type RaffleStatus string

const (
	RaffleStatusPending RaffleStatus = "pendiente"
	RaffleStatusDone    RaffleStatus = "realizado"
)

type Raffle struct {
	ID           uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	EventID      uint           `gorm:"not null;index" json:"event_id"`
	Name         string         `gorm:"type:varchar(200);not null" json:"name"`
	PricePerChance float64      `gorm:"type:decimal(10,2);not null" json:"price_per_chance"`
	Status       RaffleStatus   `gorm:"type:enum('pendiente','realizado');default:'pendiente';not null" json:"status"`
	WinnerUserID *uint          `gorm:"index" json:"winner_user_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	Event      Event  `gorm:"foreignKey:EventID" json:"event,omitempty"`
	WinnerUser *User  `gorm:"foreignKey:WinnerUserID" json:"winner_user,omitempty"`
}

type RaffleEntry struct {
	ID       uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	RaffleID uint      `gorm:"not null;index" json:"raffle_id"`
	UserID   uint      `gorm:"not null;index" json:"user_id"`
	Chances  int       `gorm:"not null" json:"chances"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Associations
	Raffle Raffle `gorm:"foreignKey:RaffleID" json:"raffle,omitempty"`
	User   User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
