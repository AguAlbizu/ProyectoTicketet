package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

// EventDAO handles all database operations for the Event model.
type EventDAO struct {
	db *gorm.DB
}

// NewEventDAO creates a new EventDAO with the provided GORM instance.
func NewEventDAO(db *gorm.DB) *EventDAO {
	return &EventDAO{db: db}
}

// Create persists a new event record.
func (d *EventDAO) Create(event *domain.Event) error {
	// TODO: d.db.Create(event)
	return nil
}

// FindAll returns all active events, optionally filtered by category.
func (d *EventDAO) FindAll(category string) ([]domain.Event, error) {
	// TODO: build query with optional WHERE category = ? filter
	// TODO: d.db.Where("status = ?", domain.EventStatusActive).Find(&events)
	return nil, nil
}

// FindByID returns a single event by its ID.
func (d *EventDAO) FindByID(id uint) (*domain.Event, error) {
	// TODO: d.db.First(&event, id)
	return nil, nil
}

// Update saves changes to an existing event.
func (d *EventDAO) Update(event *domain.Event) error {
	// TODO: d.db.Save(event)
	return nil
}

// Delete soft-deletes an event by ID.
func (d *EventDAO) Delete(id uint) error {
	// TODO: d.db.Delete(&domain.Event{}, id)
	return nil
}
