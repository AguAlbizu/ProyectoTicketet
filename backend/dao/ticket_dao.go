package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

// TicketDAO handles all database operations for the Ticket model.
type TicketDAO struct {
	db *gorm.DB
}

// NewTicketDAO creates a new TicketDAO with the provided GORM instance.
func NewTicketDAO(db *gorm.DB) *TicketDAO {
	return &TicketDAO{db: db}
}

// Create persists a new ticket record.
func (d *TicketDAO) Create(ticket *domain.Ticket) error {
	// TODO: d.db.Create(ticket)
	return nil
}

// FindByUserID returns all tickets belonging to the given user.
func (d *TicketDAO) FindByUserID(userID uint) ([]domain.Ticket, error) {
	// TODO: d.db.Preload("Event").Where("user_id = ?", userID).Find(&tickets)
	return nil, nil
}

// FindByID returns a single ticket by its ID.
func (d *TicketDAO) FindByID(id uint) (*domain.Ticket, error) {
	// TODO: d.db.Preload("Event").Preload("User").First(&ticket, id)
	return nil, nil
}

// Update saves changes to an existing ticket (e.g. status change or transfer).
func (d *TicketDAO) Update(ticket *domain.Ticket) error {
	// TODO: d.db.Save(ticket)
	return nil
}

// CountByEvent returns the number of active tickets sold for a given event.
func (d *TicketDAO) CountByEvent(eventID uint) (int64, error) {
	// TODO: d.db.Model(&domain.Ticket{}).Where("event_id = ? AND status = ?", eventID, domain.TicketStatusActive).Count(&count)
	return 0, nil
}
