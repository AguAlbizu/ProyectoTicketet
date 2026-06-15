package tests

// Objetivo de cobertura para la entrega parcial: >= 40% en servicios y controladores.
// Correr con: go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils -v

import (
	"testing"
	"ticketapp/domain"
	"ticketapp/services"

	"github.com/stretchr/testify/assert"
)

// --- Mocks ---

type mockTicketDAO struct {
	ticket    *domain.Ticket
	tickets   []domain.Ticket
	createErr error
	updateErr error
	getErr    error
	lastSaved *domain.Ticket
}

func (m *mockTicketDAO) CreateTicket(t *domain.Ticket) error {
	t.IDTickets = 1
	m.ticket = t
	return m.createErr
}
func (m *mockTicketDAO) GetTicketsByUserID(userID uint) ([]domain.Ticket, error) {
	return m.tickets, nil
}
func (m *mockTicketDAO) GetTicketByID(id uint) (*domain.Ticket, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.ticket, nil
}
func (m *mockTicketDAO) UpdateTicket(t *domain.Ticket) error {
	m.lastSaved = t
	return m.updateErr
}
func (m *mockTicketDAO) GetActiveTicketByUserAndEvent(userID, eventID uint) (*domain.Ticket, error) {
	return nil, assert.AnError
}

type mockEventDAOForTicket struct {
	event     *domain.Event
	updateErr error
	lastSaved *domain.Event
}

func (m *mockEventDAOForTicket) GetEventByID(id uint) (*domain.Event, error) {
	if m.event == nil {
		return nil, assert.AnError
	}
	return m.event, nil
}
func (m *mockEventDAOForTicket) UpdateEvent(e *domain.Event) error {
	m.lastSaved = e
	return m.updateErr
}

type mockUserDAOForTicket struct {
	user *domain.User
}

func (m *mockUserDAOForTicket) GetUserByEmail(email string) (*domain.User, error) {
	if m.user == nil {
		return nil, assert.AnError
	}
	return m.user, nil
}

// --- Tests ---

// TestBuyTicket_NoCapacity verifica que comprar cuando cupo_disponible == 0 retorna error.
func TestBuyTicket_NoCapacity(t *testing.T) {
	eventDAO := &mockEventDAOForTicket{
		event: &domain.Event{IDEvents: 1, Titulo: "Evento lleno", Estado: "activo", CupoDisponible: 0},
	}
	svc := services.NewTicketService(&mockTicketDAO{}, eventDAO, &mockUserDAOForTicket{})

	_, err := svc.BuyTicket(1, 1, false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no hay cupo disponible")
}

// TestCancelTicket_NotOwner verifica que cancelar un ticket ajeno retorna error 403.
func TestCancelTicket_NotOwner(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		ticket: &domain.Ticket{IDTickets: 1, IDUsers: 99, Estado: "activo"},
	}
	svc := services.NewTicketService(ticketDAO, &mockEventDAOForTicket{event: &domain.Event{IDEvents: 1}}, &mockUserDAOForTicket{})

	err := svc.CancelTicket(1, 1) // usuario 1 intenta cancelar ticket del usuario 99
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no tenés permiso")
}

// TestCancelTicket_RestoresCapacity verifica que al cancelar el cupo del evento se incrementa en 1.
func TestCancelTicket_RestoresCapacity(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		ticket: &domain.Ticket{IDTickets: 1, IDUsers: 1, IDEvents: 1, Estado: "activo"},
	}
	eventDAO := &mockEventDAOForTicket{
		event: &domain.Event{IDEvents: 1, CupoDisponible: 5},
	}
	svc := services.NewTicketService(ticketDAO, eventDAO, &mockUserDAOForTicket{})

	err := svc.CancelTicket(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, "cancelado", ticketDAO.lastSaved.Estado)
	assert.Equal(t, 6, eventDAO.lastSaved.CupoDisponible)
}

// TestBuyTicket_Success verifica que comprar con cupo disponible crea el ticket y decrementa el cupo.
func TestBuyTicket_Success(t *testing.T) {
	ticketDAO := &mockTicketDAO{}
	eventDAO := &mockEventDAOForTicket{
		event: &domain.Event{IDEvents: 1, Titulo: "Concierto", Estado: "activo", CupoDisponible: 10},
	}
	svc := services.NewTicketService(ticketDAO, eventDAO, &mockUserDAOForTicket{})

	ticket, err := svc.BuyTicket(1, 1, false)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, 9, eventDAO.lastSaved.CupoDisponible)
}

// TestBuyTicket_EventCancelled verifica que comprar en un evento cancelado retorna error.
func TestBuyTicket_EventCancelled(t *testing.T) {
	eventDAO := &mockEventDAOForTicket{
		event: &domain.Event{IDEvents: 1, Estado: "cancelado", CupoDisponible: 5},
	}
	svc := services.NewTicketService(&mockTicketDAO{}, eventDAO, &mockUserDAOForTicket{})

	_, err := svc.BuyTicket(1, 1, false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "el evento está cancelado")
}

// TestGetMyTickets_ReturnsTickets verifica que retorna la lista de tickets del usuario.
func TestGetMyTickets_ReturnsTickets(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		tickets: []domain.Ticket{
			{IDTickets: 1, IDUsers: 1, IDEvents: 1, Estado: "activo"},
			{IDTickets: 2, IDUsers: 1, IDEvents: 2, Estado: "cancelado"},
		},
	}
	svc := services.NewTicketService(ticketDAO, &mockEventDAOForTicket{}, &mockUserDAOForTicket{})

	tickets, err := svc.GetMyTickets(1)
	assert.NoError(t, err)
	assert.Len(t, tickets, 2)
}

// TestTransferTicket_Success verifica que una transferencia exitosa cambia el estado a "transferido".
func TestTransferTicket_Success(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		ticket: &domain.Ticket{IDTickets: 1, IDUsers: 1, IDEvents: 1, Estado: "activo"},
	}
	userDAO := &mockUserDAOForTicket{
		user: &domain.User{IDUsers: 2, Email: "destino@test.com"},
	}
	svc := services.NewTicketService(ticketDAO, &mockEventDAOForTicket{event: &domain.Event{IDEvents: 1}}, userDAO)

	err := svc.TransferTicket(1, 1, "destino@test.com")
	assert.NoError(t, err)
	assert.Equal(t, "transferido", ticketDAO.lastSaved.Estado)
}

// TestTransferTicket_NotOwner verifica que transferir un ticket ajeno retorna error 403.
func TestTransferTicket_NotOwner(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		ticket: &domain.Ticket{IDTickets: 1, IDUsers: 99, Estado: "activo"},
	}
	svc := services.NewTicketService(ticketDAO, &mockEventDAOForTicket{event: &domain.Event{}}, &mockUserDAOForTicket{})

	err := svc.TransferTicket(1, 1, "cualquiera@test.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no tenés permiso")
}

// TestTransferTicket_NotActive verifica que transferir un ticket inactivo retorna error.
func TestTransferTicket_NotActive(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		ticket: &domain.Ticket{IDTickets: 1, IDUsers: 1, Estado: "cancelado"},
	}
	svc := services.NewTicketService(ticketDAO, &mockEventDAOForTicket{event: &domain.Event{}}, &mockUserDAOForTicket{})

	err := svc.TransferTicket(1, 1, "cualquiera@test.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "solo se pueden transferir entradas activas")
}

// TestTransferTicket_TargetNotFound verifica que transferir a un email inexistente retorna error.
func TestTransferTicket_TargetNotFound(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		ticket: &domain.Ticket{IDTickets: 1, IDUsers: 1, IDEvents: 1, Estado: "activo"},
	}
	svc := services.NewTicketService(ticketDAO, &mockEventDAOForTicket{event: &domain.Event{IDEvents: 1}}, &mockUserDAOForTicket{})

	err := svc.TransferTicket(1, 1, "noexiste@test.com")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "usuario destino no encontrado")
}
