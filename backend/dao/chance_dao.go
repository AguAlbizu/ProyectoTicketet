package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

type ChanceDAO struct {
	db *gorm.DB
}

func NewChanceDAO(db *gorm.DB) *ChanceDAO {
	return &ChanceDAO{db: db}
}

func (d *ChanceDAO) CreateChance(chance *domain.Chance) error {
	return d.db.Create(chance).Error
}

// GetChancesBySorteoID retorna todas las chances del sorteo con el usuario precargado.
// Cada fila representa una chance individual (a mayor cantidad de filas, mayor probabilidad de ganar).
func (d *ChanceDAO) GetChancesBySorteoID(sorteoID uint) ([]domain.Chance, error) {
	var chances []domain.Chance
	err := d.db.Preload("User").Where("id_sorteo = ?", sorteoID).Find(&chances).Error
	return chances, err
}

func (d *ChanceDAO) CountChancesByUserAndSorteo(userID, sorteoID uint) (int64, error) {
	var count int64
	err := d.db.Model(&domain.Chance{}).Where("id_sorteo = ? AND id_users = ?", sorteoID, userID).Count(&count).Error
	return count, err
}
