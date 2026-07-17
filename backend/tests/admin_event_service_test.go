package tests

// Tests unitarios de AdminEventService, con foco en GetEventReport
// (separación de compradores originales vs. titulares con entrada activa).

import (
	"testing"
	"ticketapp/domain"
	"ticketapp/services"
	"time"

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
	tickets        []domain.Ticket
	getErr         error
	cancelAllErr   error
	cancelAllCalls int
}

func (m *mockAdminTicketDAO) GetTicketsByEventID(eventID uint) ([]domain.Ticket, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.tickets, nil
}
func (m *mockAdminTicketDAO) CancelAllTicketsByEventID(eventID uint) error {
	m.cancelAllCalls++
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

// --- Tests: GetAllEvents ---

func TestGetAllEvents_ReturnsEvents(t *testing.T) {
	eventDAO := &mockAdminEventDAO{events: []domain.Event{
		{IDEvents: 1, Titulo: "Evento A"},
		{IDEvents: 2, Titulo: "Evento B"},
	}}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	events, err := svc.GetAllEvents()
	assert.NoError(t, err)
	assert.Len(t, events, 2)
}

// --- Tests: CreateEvent ---

func TestCreateEvent_Success(t *testing.T) {
	svc := services.NewAdminEventService(&mockAdminEventDAO{}, &mockAdminTicketDAO{})

	event, err := svc.CreateEvent(services.CreateEventInput{
		Titulo: "Recital", Fecha: time.Now(), Hora: "20:00", Capacidad: 100,
	})
	assert.NoError(t, err)
	assert.Equal(t, "activo", event.Estado)
	assert.Equal(t, 100, event.CupoDisponible, "el cupo inicial debe ser igual a la capacidad")
}

func TestCreateEvent_DAOError(t *testing.T) {
	eventDAO := &mockAdminEventDAO{createErr: assert.AnError}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	_, err := svc.CreateEvent(services.CreateEventInput{Titulo: "Recital", Capacidad: 100})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error al crear el evento")
}

// --- Tests: UpdateEvent ---

func TestUpdateEvent_NotFound(t *testing.T) {
	eventDAO := &mockAdminEventDAO{getErr: assert.AnError}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	_, err := svc.UpdateEvent(99, services.UpdateEventInput{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "evento no encontrado")
}

func TestUpdateEvent_Deactivate_CancelsTickets(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo", Capacidad: 100, CupoDisponible: 80}}
	ticketDAO := &mockAdminTicketDAO{}
	svc := services.NewAdminEventService(eventDAO, ticketDAO)

	event, err := svc.UpdateEvent(1, services.UpdateEventInput{Titulo: "X", Capacidad: 100, Estado: "cancelado"})
	assert.NoError(t, err)
	assert.Equal(t, "cancelado", event.Estado)
	assert.Equal(t, 1, ticketDAO.cancelAllCalls, "debe cancelar todas las entradas al desactivar el evento")
}

func TestUpdateEvent_Reactivate_ResetsCupo(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "cancelado", Capacidad: 50, CupoDisponible: 50}}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	event, err := svc.UpdateEvent(1, services.UpdateEventInput{Titulo: "X", Capacidad: 80, Estado: "activo"})
	assert.NoError(t, err)
	assert.Equal(t, "activo", event.Estado)
	assert.Equal(t, 80, event.CupoDisponible)
}

func TestUpdateEvent_NoStatusChange_RecalculatesCupo(t *testing.T) {
	// Capacidad actual 100, cupo disponible 80 => 20 vendidas. Nueva capacidad 50 => cupo 30.
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo", Capacidad: 100, CupoDisponible: 80}}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	event, err := svc.UpdateEvent(1, services.UpdateEventInput{Titulo: "X", Capacidad: 50})
	assert.NoError(t, err)
	assert.Equal(t, 30, event.CupoDisponible)
}

func TestUpdateEvent_NewCapacityBelowSold_ReturnsError(t *testing.T) {
	// 20 vendidas, nueva capacidad 10 no alcanza.
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo", Capacidad: 100, CupoDisponible: 80}}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	_, err := svc.UpdateEvent(1, services.UpdateEventInput{Titulo: "X", Capacidad: 10})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "menor a las entradas ya vendidas")
}

func TestUpdateEvent_CancelTicketsError(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo", Capacidad: 100, CupoDisponible: 80}}
	ticketDAO := &mockAdminTicketDAO{cancelAllErr: assert.AnError}
	svc := services.NewAdminEventService(eventDAO, ticketDAO)

	_, err := svc.UpdateEvent(1, services.UpdateEventInput{Titulo: "X", Capacidad: 100, Estado: "cancelado"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error al cancelar las entradas")
}

func TestUpdateEvent_FullUpdateError(t *testing.T) {
	eventDAO := &mockAdminEventDAO{
		event:         &domain.Event{IDEvents: 1, Estado: "activo", Capacidad: 100, CupoDisponible: 80},
		fullUpdateErr: assert.AnError,
	}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	_, err := svc.UpdateEvent(1, services.UpdateEventInput{Titulo: "X", Capacidad: 100})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error al actualizar el evento")
}

// --- Tests: CancelEvent ---

func TestCancelEvent_Success(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo"}}
	ticketDAO := &mockAdminTicketDAO{}
	svc := services.NewAdminEventService(eventDAO, ticketDAO)

	err := svc.CancelEvent(1)
	assert.NoError(t, err)
	assert.Equal(t, 1, ticketDAO.cancelAllCalls)
}

func TestCancelEvent_NotFound(t *testing.T) {
	eventDAO := &mockAdminEventDAO{getErr: assert.AnError}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	err := svc.CancelEvent(99)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "evento no encontrado")
}

func TestCancelEvent_AlreadyCancelled(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "cancelado"}}
	svc := services.NewAdminEventService(eventDAO, &mockAdminTicketDAO{})

	err := svc.CancelEvent(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ya está cancelado")
}

func TestCancelEvent_CancelTicketsError(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo"}}
	ticketDAO := &mockAdminTicketDAO{cancelAllErr: assert.AnError}
	svc := services.NewAdminEventService(eventDAO, ticketDAO)

	err := svc.CancelEvent(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error al cancelar las entradas")
}
