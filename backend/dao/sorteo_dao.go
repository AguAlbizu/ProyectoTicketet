package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

type SorteoDAO struct {
	db *gorm.DB
}

func NewSorteoDAO(db *gorm.DB) *SorteoDAO {
	return &SorteoDAO{db: db}
}

func (d *SorteoDAO) CreateSorteo(sorteo *domain.Sorteo) error {
	return d.db.Create(sorteo).Error
}

func (d *SorteoDAO) GetSorteoByID(id uint) (*domain.Sorteo, error) {
	var sorteo domain.Sorteo
	err := d.db.First(&sorteo, id).Error
	if err != nil {
		return nil, err
	}
	return &sorteo, nil
}

func (d *SorteoDAO) GetSorteoByEventID(eventID uint) (*domain.Sorteo, error) {
	var sorteo domain.Sorteo
	err := d.db.Where("id_events = ?", eventID).First(&sorteo).Error
	if err != nil {
		return nil, err
	}
	return &sorteo, nil
}

func (d *SorteoDAO) UpdateSorteo(sorteo *domain.Sorteo) error {
	return d.db.Exec(
		"UPDATE sorteos SET estado = ?, id_ganador = ?, fecha_realizado = ? WHERE id_sorteo = ?",
		sorteo.Estado, sorteo.IDGanador, sorteo.FechaRealizado, sorteo.IDSorteo,
	).Error
}

// GetSorteosConEvento retorna todos los sorteos con su evento y ganador precargados,
// usado por el panel de administración para listar eventos con sorteo.
func (d *SorteoDAO) GetSorteosConEvento() ([]domain.Sorteo, error) {
	var sorteos []domain.Sorteo
	err := d.db.Preload("Event").Preload("Ganador").Order("created_at desc").Find(&sorteos).Error
	return sorteos, err
}
