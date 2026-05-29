package domain

import "time"

type Ticket struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	User        User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	EventID     uint      `gorm:"not null;index" json:"event_id"`
	Event       Event     `gorm:"foreignKey:EventID" json:"event,omitempty"`
	Estado      string    `gorm:"type:varchar(20);default:'activo';not null" json:"estado"`
	FechaCompra time.Time `gorm:"not null" json:"fecha_compra"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Ticket) TableName() string {
	return "tickets"
}
