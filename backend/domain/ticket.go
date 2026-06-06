package domain

import "time"

type Ticket struct {
	IDTickets   uint      `gorm:"primaryKey;autoIncrement;column:id_tickets" json:"id_tickets"`
	IDUsers     uint      `gorm:"not null;index;column:id_users" json:"id_users"`
	User        User      `gorm:"foreignKey:IDUsers;references:IDUsers" json:"user,omitempty"`
	IDEvents    uint      `gorm:"not null;index;column:id_events" json:"id_events"`
	Event       Event     `gorm:"foreignKey:IDEvents;references:IDEvents" json:"event,omitempty"`
	Estado      string    `gorm:"type:varchar(20);default:'activo';not null" json:"estado"`
	FechaCompra time.Time `gorm:"not null" json:"fecha_compra"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Ticket) TableName() string {
	return "tickets"
}
