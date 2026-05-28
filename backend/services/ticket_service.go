package services

import (
	"ticketapp/dao"
	"ticketapp/domain"
)

// TicketService handles business logic for ticket purchasing, cancellation, and transfer.
type TicketService struct {
	ticketDAO *dao.TicketDAO
	eventDAO  *dao.EventDAO
	userDAO   *dao.UserDAO
}

// NewTicketService creates a new TicketService with its required dependencies.
func NewTicketService(ticketDAO *dao.TicketDAO, eventDAO *dao.EventDAO, userDAO *dao.UserDAO) *TicketService {
	return &TicketService{
		ticketDAO: ticketDAO,
		eventDAO:  eventDAO,
		userDAO:   userDAO,
	}
}

// Purchase creates a ticket for a user for a given event.
// Validates that the event is active and has available tickets, then decrements AvailableTickets.
func (s *TicketService) Purchase(userID, eventID uint) (*domain.Ticket, error) {
	// TODO: fetch event with eventDAO.FindByID, verify status == active and available_tickets > 0
	// TODO: create domain.Ticket{UserID, EventID, Status: active, PurchaseDate: now}
	// TODO: call ticketDAO.Create
	// TODO: decrement event.AvailableTickets and call eventDAO.Update
	return nil, nil
}

// GetByUser returns all tickets belonging to the authenticated user.
func (s *TicketService) GetByUser(userID uint) ([]domain.Ticket, error) {
	// TODO: delegate to ticketDAO.FindByUserID(userID)
	return nil, nil
}

// Cancel marks a ticket as cancelled and restores the event's available slot.
func (s *TicketService) Cancel(ticketID, userID uint) error {
	// TODO: fetch ticket, verify it belongs to userID and status == active
	// TODO: set ticket.Status = cancelled, call ticketDAO.Update
	// TODO: increment event.AvailableTickets, call eventDAO.Update
	return nil
}

// Transfer reassigns a ticket to a different user by email.
func (s *TicketService) Transfer(ticketID, ownerID uint, targetEmail string) error {
	// TODO: fetch ticket, verify it belongs to ownerID and status == active
	// TODO: find target user by email with userDAO.FindByEmail
	// TODO: set ticket.UserID = target.ID, Status = transferred, call ticketDAO.Update
	return nil
}
