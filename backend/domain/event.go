package domain

import "time"

type Event struct {
	ID             uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Titulo         string    `gorm:"type:varchar(200);not null" json:"titulo"`
	Descripcion    string    `gorm:"type:text" json:"descripcion"`
	Fecha          time.Time `gorm:"not null" json:"fecha"`
	Hora           string    `gorm:"type:varchar(5);not null" json:"hora"`
	Duracion       string    `gorm:"type:varchar(50)" json:"duracion"`
	Capacidad      int       `gorm:"not null" json:"capacidad"`
	CupoDisponible int       `gorm:"not null" json:"cupo_disponible"`
	Categoria      string    `gorm:"type:varchar(100)" json:"categoria"`
	ImagenURL      string    `gorm:"type:varchar(500)" json:"imagen_url"`
	Estado         string    `gorm:"type:varchar(20);default:'activo';not null" json:"estado"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Event) TableName() string {
	return "events"
}
