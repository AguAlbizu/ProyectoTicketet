package services

import (
	"ticketapp/dao"
	"ticketapp/domain"
)

// EventService handles business logic for event management.
type EventService struct {
	eventDAO *dao.EventDAO
}

// NewEventService creates a new EventService with its required dependencies.
func NewEventService(eventDAO *dao.EventDAO) *EventService {
	return &EventService{eventDAO: eventDAO}
}

// EventInput holds the data required to create or update an event.
type EventInput struct {
	Title            string
	Description      string
	Date             string
	DurationMinutes  int
	Capacity         int
	Category         string
	ImageURL         string
}

// GetAll returns all active events, optionally filtered by category.
func (s *EventService) GetAll(category string) ([]domain.Event, error) {
	// TODO: delegate to eventDAO.FindAll(category)
	return nil, nil
}

// GetByID returns a single event by ID or an error if not found.
func (s *EventService) GetByID(id uint) (*domain.Event, error) {
	// TODO: delegate to eventDAO.FindByID(id)
	return nil, nil
}

// Create validates and creates a new event (admin only).
// Sets AvailableTickets = Capacity on creation.
func (s *EventService) Create(input EventInput) (*domain.Event, error) {
	// TODO: validate required fields (title, date, capacity > 0)
	// TODO: parse date string to time.Time
	// TODO: build domain.Event and call eventDAO.Create
	return nil, nil
}

// Update applies changes to an existing event (admin only).
func (s *EventService) Update(id uint, input EventInput) (*domain.Event, error) {
	// TODO: fetch event with eventDAO.FindByID
	// TODO: apply changes from input to fetched event
	// TODO: call eventDAO.Update
	return nil, nil
}

// Cancel changes an event status to cancelled.
func (s *EventService) Cancel(id uint) error {
	// TODO: fetch event, set Status = EventStatusCancelled, call eventDAO.Update
	return nil
}
