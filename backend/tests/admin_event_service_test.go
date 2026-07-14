package tests

// Tests unitarios de AdminEventService, con foco en GetEventReport
// (separación de compradores originales vs. titulares con entrada activa).

import (
	"testing"
	"ticketapp/domain"
	"ticketapp/services"

	"github.com/stretchr/testify/assert"
)

// --- Mocks ---

type mockAdminEventDAO struct {
	events        []domain.Event
	event         *domain.Event
	getErr        error
	createErr     error
	fullUpdateErr error
}

func (m *mockAdminEventDAO) GetAllEventsAdmin() ([]domain.Event, error) {
	return m.events, nil
}
func (m *mockAdminEventDAO) GetEventByID(id uint) (*domain.Event, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.event, nil
}
func (m *mockAdminEventDAO) CreateEvent(event *domain.Event) error {
	return m.createErr
}
func (m *mockAdminEventDAO) FullUpdateEvent(event *domain.Event) error {
	return m.fullUpdateErr
}

type mockAdminTicketDAO struct {
	tickets      []domain.Ticket
	getErr       error
	cancelAllErr error
}

func (m *mockAdminTicketDAO) GetTicketsByEventID(eventID uint) ([]domain.Ticket, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.tickets, nil
}
func (m *mockAdminTicketDAO) CancelAllTicketsByEventID(eventID uint) error {
	return m.cancelAllErr
}

// --- Tests: GetEventReport ---

func TestGetEventReport_SeparaCompradoresYActivos(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Titulo: "Rifa", Capacidad: 100, CupoDisponible: 97}}
	ticketDAO := &mockAdminTicketDAO{tickets: []domain.Ticket{
		// Compró y sigue activa
		{IDTickets: 1, IDUsers: 1, User: domain.User{Nombre: "Ana"}, Estado: "activo", Origen: "compra"},
		// Compró y la canceló
		{IDTickets: 2, IDUsers: 2, User: domain.User{Nombre: "Beto"}, Estado: "cancelado", Origen: "compra"},
		// Compró y la transfirió (ya no la tiene)
		{IDTickets: 3, IDUsers: 3, User: domain.User{Nombre: "Caro"}, Estado: "transferido", Origen: "compra"},
		// La recibió por transferencia, la tiene activa (no es "comprador" original)
		{IDTickets: 4, IDUsers: 4, User: domain.User{Nombre: "Dana"}, Estado: "activo", Origen: "transferencia"},
	}}
	svc := services.NewAdminEventService(eventDAO, ticketDAO)

	report, err := svc.GetEventReport(1)
	assert.NoError(t, err)

	// Compradores: solo origen "compra", en cualquier estado
	assert.Len(t, report.Compradores, 3)
	var beto *services.BuyerInfo
	for i := range report.Compradores {
		if report.Compradores[i].Nombre == "Beto" {
			beto = &report.Compradores[i]
		}
	}
	assert.NotNil(t, beto, "Beto canceló pero debe seguir apareciendo en compradores")
	assert.Equal(t, "cancelado", beto.Estado)

	// Titulares activos: cualquier origen, solo estado activo
	assert.Len(t, report.TitularesActivos, 2)
	nombres := []string{report.TitularesActivos[0].Nombre, report.TitularesActivos[1].Nombre}
	assert.Contains(t, nombres, "Ana")
	assert.Contains(t, nombres, "Dana")
}

func TestGetEventReport_EventNotFound(t *testing.T) {
	eventDAO := &mockAdminEventDAO{getErr: assert.AnError}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	_, err := svc.GetEventReport(99)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "evento no encontrado")
}
