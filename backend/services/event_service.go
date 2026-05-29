package services

import (
	"fmt"
	"ticketapp/domain"
)

// EventRepository define los métodos de persistencia requeridos por EventService.
type EventRepository interface {
	GetAllEvents(categoria string) ([]domain.Event, error)
	GetEventByID(id uint) (*domain.Event, error)
}

type EventService struct {
	eventDAO EventRepository
}

func NewEventService(eventDAO EventRepository) *EventService {
	return &EventService{eventDAO: eventDAO}
}

// GetEvents retorna los eventos activos. Si categoria no es vacía, filtra por ella.
func (s *EventService) GetEvents(categoria string) ([]domain.Event, error) {
	return s.eventDAO.GetAllEvents(categoria)
}

// GetEventByID retorna el evento o error si no existe o está cancelado.
func (s *EventService) GetEventByID(id uint) (*domain.Event, error) {
	event, err := s.eventDAO.GetEventByID(id)
	if err != nil {
		return nil, fmt.Errorf("evento no encontrado")
	}
	if event.Estado == "cancelado" {
		return nil, fmt.Errorf("el evento está cancelado")
	}
	return event, nil
}
