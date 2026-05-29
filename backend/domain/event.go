package domain

import (
	"time"
)

type Event struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title            string    `gorm:"type:varchar(200);not null" json:"titulo"`
	Description      string    `gorm:"type:text" json:"descripcion"`
	Date             time.Time `gorm:"not null" json:"fecha"`
	Time             string    `gorm:"type:varchar(5);not null" json:"hora"`
	Duration         int       `gorm:"not null" json:"duracion_minutos"`
	Capacity         int       `gorm:"not null" json:"capacidad"`
	AvailableTickets int       `gorm:"not null" json:"cupo_disponible"`
	Category         string    `gorm:"type:varchar(100)" json:"categoria"`
	ImageURL         string    `gorm:"type:varchar(500)" json:"imagen_url"`
	Status           string    `gorm:"type:varchar(20);default:'activo';not null" json:"estado"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
