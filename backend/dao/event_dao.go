package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

type EventDAO struct {
	db *gorm.DB
}

func NewEventDAO(db *gorm.DB) *EventDAO {
	return &EventDAO{db: db}
}

// GetAllEvents retorna todos los eventos activos. Si categoria no es vacía, filtra por ella.
func (d *EventDAO) GetAllEvents(categoria string) ([]domain.Event, error) {
	var events []domain.Event
	query := d.db.Where("estado = ?", "activo")
	if categoria != "" {
		query = query.Where("categoria = ?", categoria)
	}
	err := query.Find(&events).Error
	return events, err
}

func (d *EventDAO) GetEventByID(id uint) (*domain.Event, error) {
	var event domain.Event
	err := d.db.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (d *EventDAO) UpdateEvent(event *domain.Event) error {
	return d.db.Exec(
		"UPDATE events SET cupo_disponible = ?, estado = ? WHERE id_events = ?",
		event.CupoDisponible, event.Estado, event.IDEvents,
	).Error
}

func (d *EventDAO) GetAllEventsAdmin() ([]domain.Event, error) {
	var events []domain.Event
	err := d.db.Order("created_at DESC").Find(&events).Error
	return events, err
}

func (d *EventDAO) CreateEvent(event *domain.Event) error {
	return d.db.Create(event).Error
}

func (d *EventDAO) FullUpdateEvent(event *domain.Event) error {
	return d.db.Exec(
		`UPDATE events SET titulo=?, descripcion=?, fecha=?, hora=?, capacidad=?,
		 cupo_disponible=?, categoria=?, direccion=?, imagen_url=?, precio=?, estado=?
		 WHERE id_events=?`,
		event.Titulo, event.Descripcion, event.Fecha, event.Hora, event.Capacidad,
		event.CupoDisponible, event.Categoria, event.Direccion, event.ImagenURL,
		event.Precio, event.Estado, event.IDEvents,
	).Error
}
