package domain

import "time"

// TODO (entrega final): implementar validación de roles en middleware JWT

type User struct {
	IDUsers   uint      `gorm:"primaryKey;autoIncrement;column:id_users" json:"id_users"`
	Nombre    string    `gorm:"type:varchar(100);not null" json:"nombre"`
	Email     string    `gorm:"type:varchar(150);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Rol       string    `gorm:"type:varchar(20);default:'cliente';not null" json:"rol"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
