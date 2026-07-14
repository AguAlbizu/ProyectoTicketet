package domain

import "time"

// Notification representa una notificación in-app para un usuario
// (ej. confirmación de participación en un sorteo, resultado del sorteo).
type Notification struct {
	IDNotification uint      `gorm:"primaryKey;autoIncrement;column:id_notification" json:"id_notification"`
	IDUsers        uint      `gorm:"not null;index;column:id_users" json:"id_users"`
	Tipo           string    `gorm:"type:varchar(30);not null" json:"tipo"`
	Titulo         string    `gorm:"type:varchar(150);not null" json:"titulo"`
	Mensaje        string    `gorm:"type:text;not null" json:"mensaje"`
	Leida          bool      `gorm:"default:false;not null" json:"leida"`
	IDSorteo       *uint     `gorm:"column:id_sorteo" json:"id_sorteo"`
	CreatedAt      time.Time `json:"created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
