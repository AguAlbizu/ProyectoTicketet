package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

type TicketDAO struct {
	db *gorm.DB
}

func NewTicketDAO(db *gorm.DB) *TicketDAO {
	return &TicketDAO{db: db}
}

func (d *TicketDAO) CreateTicket(ticket *domain.Ticket) error {
	return d.db.Create(ticket).Error
}

// GetTicketsByUserID retorna todos los tickets del usuario con el evento precargado.
func (d *TicketDAO) GetTicketsByUserID(userID uint) ([]domain.Ticket, error) {
	var tickets []domain.Ticket
	err := d.db.Preload("Event").Where("id_users = ?", userID).Find(&tickets).Error
	return tickets, err
}

func (d *TicketDAO) GetTicketByID(id uint) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := d.db.Preload("Event").First(&ticket, id).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (d *TicketDAO) UpdateTicket(ticket *domain.Ticket) error {
	return d.db.Exec(
		"UPDATE tickets SET estado = ? WHERE id_tickets = ?",
		ticket.Estado, ticket.IDTickets,
	).Error
}

func (d *TicketDAO) GetActiveTicketByUserAndEvent(userID, eventID uint) (*domain.Ticket, error) {
	var ticket domain.Ticket
	err := d.db.Where("id_users = ? AND id_events = ? AND estado = 'activo'", userID, eventID).First(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (d *TicketDAO) GetTicketsByEventID(eventID uint) ([]domain.Ticket, error) {
	var tickets []domain.Ticket
	err := d.db.Preload("User").Where("id_events = ? AND estado = 'activo'", eventID).Find(&tickets).Error
	return tickets, err
}

func (d *TicketDAO) CancelAllTicketsByEventID(eventID uint) error {
	return d.db.Exec(
		"UPDATE tickets SET estado = 'cancelado' WHERE id_events = ? AND estado = 'activo'",
		eventID,
	).Error
}
