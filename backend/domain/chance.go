package domain

import "time"

// Chance representa una chance individual comprada por un usuario para un sorteo.
// Cada fila es una entrada al sorteo; comprar varias chances aumenta las probabilidades de ganar.
type Chance struct {
	IDChance    uint      `gorm:"primaryKey;autoIncrement;column:id_chance" json:"id_chance"`
	IDSorteo    uint      `gorm:"not null;index;column:id_sorteo" json:"id_sorteo"`
	IDUsers     uint      `gorm:"not null;index;column:id_users" json:"id_users"`
	User        User      `gorm:"foreignKey:IDUsers;references:IDUsers" json:"user,omitempty"`
	FechaCompra time.Time `gorm:"not null" json:"fecha_compra"`
	CreatedAt   time.Time `json:"created_at"`
}

func (Chance) TableName() string {
	return "chances"
}
