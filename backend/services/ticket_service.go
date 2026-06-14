package services

import (
	"fmt"
	"time"
	"ticketapp/domain"
)

// TicketDAOPort define los métodos de persistencia de tickets.
type TicketDAOPort interface {
	CreateTicket(ticket *domain.Ticket) error
	GetTicketsByUserID(userID uint) ([]domain.Ticket, error)
	GetTicketByID(id uint) (*domain.Ticket, error)
	UpdateTicket(ticket *domain.Ticket) error
}

// EventDAOPort define los métodos de eventos requeridos por TicketService.
type EventDAOPort interface {
	GetEventByID(id uint) (*domain.Event, error)
	UpdateEvent(event *domain.Event) error
}

// UserDAOPort define los métodos de usuarios requeridos por TicketService.
type UserDAOPort interface {
	GetUserByEmail(email string) (*domain.User, error)
}

type TicketService struct {
	ticketDAO TicketDAOPort
	eventDAO  EventDAOPort
	userDAO   UserDAOPort
}

func NewTicketService(ticketDAO TicketDAOPort, eventDAO EventDAOPort, userDAO UserDAOPort) *TicketService {
	return &TicketService{ticketDAO: ticketDAO, eventDAO: eventDAO, userDAO: userDAO}
}

// BuyTicket verifica cupo, crea el ticket y decrementa CupoDisponible del evento.
func (s *TicketService) BuyTicket(userID, eventID uint) (*domain.Ticket, error) {
	event, err := s.eventDAO.GetEventByID(eventID)
	if err != nil {
		return nil, fmt.Errorf("evento no encontrado")
	}
	if event.Estado == "cancelado" {
		return nil, fmt.Errorf("el evento está cancelado")
	}
	if event.CupoDisponible <= 0 {
		return nil, fmt.Errorf("no hay cupo disponible para este evento")
	}

	ticket := &domain.Ticket{
		IDUsers:     userID,
		IDEvents:    eventID,
		Estado:      "activo",
		Origen:      "compra",
		FechaCompra: time.Now(),
	}
	if err := s.ticketDAO.CreateTicket(ticket); err != nil {
		return nil, fmt.Errorf("error al crear la entrada: %w", err)
	}

	event.CupoDisponible--
	if err := s.eventDAO.UpdateEvent(event); err != nil {
		// TODO (entrega final): revertir ticket si falla la actualización del cupo
		return nil, fmt.Errorf("error al actualizar el cupo del evento: %w", err)
	}

	return s.ticketDAO.GetTicketByID(ticket.IDTickets)
}

// GetMyTickets retorna todas las entradas del usuario con el evento precargado.
func (s *TicketService) GetMyTickets(userID uint) ([]domain.Ticket, error) {
	return s.ticketDAO.GetTicketsByUserID(userID)
}

// CancelTicket cancela una entrada activa y devuelve el cupo al evento.
func (s *TicketService) CancelTicket(ticketID, userID uint) error {
	ticket, err := s.ticketDAO.GetTicketByID(ticketID)
	if err != nil {
		return fmt.Errorf("entrada no encontrada")
	}
	if ticket.IDUsers != userID {
		return fmt.Errorf("no tenés permiso para cancelar esta entrada")
	}
	if ticket.Estado != "activo" {
		return fmt.Errorf("solo se pueden cancelar entradas activas")
	}

	ticket.Estado = "cancelado"
	if err := s.ticketDAO.UpdateTicket(ticket); err != nil {
		return fmt.Errorf("error al cancelar la entrada: %w", err)
	}

	// Devolver cupo al evento
	event, err := s.eventDAO.GetEventByID(ticket.IDEvents)
	if err == nil {
		event.CupoDisponible++
		// TODO (entrega final): manejar el error con una transacción
		s.eventDAO.UpdateEvent(event)
	}
	return nil
}

// TransferTicket transfiere una entrada activa a otro usuario y crea una nueva entrada para el destino.
func (s *TicketService) TransferTicket(ticketID, ownerID uint, targetEmail string) error {
	ticket, err := s.ticketDAO.GetTicketByID(ticketID)
	if err != nil {
		return fmt.Errorf("entrada no encontrada")
	}
	if ticket.IDUsers != ownerID {
		return fmt.Errorf("no tenés permiso para transferir esta entrada")
	}
	if ticket.Estado != "activo" {
		return fmt.Errorf("solo se pueden transferir entradas activas")
	}

	targetUser, err := s.userDAO.GetUserByEmail(targetEmail)
	if err != nil {
		return fmt.Errorf("usuario destino no encontrado")
	}

	ticket.Estado = "transferido"
	if err := s.ticketDAO.UpdateTicket(ticket); err != nil {
		return fmt.Errorf("error al transferir la entrada: %w", err)
	}

	newTicket := &domain.Ticket{
		IDUsers:     targetUser.IDUsers,
		IDEvents:    ticket.IDEvents,
		Estado:      "activo",
		Origen:      "transferencia",
		FechaCompra: time.Now(),
	}
	if err := s.ticketDAO.CreateTicket(newTicket); err != nil {
		return fmt.Errorf("error al crear la entrada para el destinatario: %w", err)
	}
	return nil
}
