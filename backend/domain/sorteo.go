package domain

import "time"

// Sorteo representa el sorteo opcional asociado a un evento (Bonus Track).
// Estado: activo (admite chances) | realizado (ya tiene ganador) | cancelado.
type Sorteo struct {
	IDSorteo       uint       `gorm:"primaryKey;autoIncrement;column:id_sorteo" json:"id_sorteo"`
	IDEvents       uint       `gorm:"not null;uniqueIndex;column:id_events" json:"id_events"`
	Event          Event      `gorm:"foreignKey:IDEvents;references:IDEvents" json:"event,omitempty"`
	Nombre         string     `gorm:"type:varchar(150);not null" json:"nombre"`
	ValorChance    int        `gorm:"not null" json:"valor_chance"`
	Estado         string     `gorm:"type:varchar(20);default:'activo';not null" json:"estado"`
	IDGanador      *uint      `gorm:"column:id_ganador" json:"id_ganador"`
	Ganador        *User      `gorm:"foreignKey:IDGanador;references:IDUsers" json:"ganador,omitempty"`
	FechaRealizado *time.Time `json:"fecha_realizado"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (Sorteo) TableName() string {
	return "sorteos"
}
