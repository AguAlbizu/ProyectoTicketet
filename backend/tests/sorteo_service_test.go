package tests

// Tests unitarios de SorteoService (funcionalidad Bonus Track).
// Correr con: go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils,ticketapp/controllers -v

import (
	"testing"
	"ticketapp/domain"
	"ticketapp/services"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// --- Mocks ---

type mockSorteoDAO struct {
	sorteo         *domain.Sorteo
	sorteos        []domain.Sorteo
	sorteosByEvent []domain.Sorteo
	activeSorteo   *domain.Sorteo
	getByIDErr     error
	getByEvtErr    error
	getActiveErr   error
	createErr      error
	updateErr      error
	lastSaved      *domain.Sorteo
}

func (m *mockSorteoDAO) CreateSorteo(s *domain.Sorteo) error {
	s.IDSorteo = 1
	m.sorteo = s
	return m.createErr
}
func (m *mockSorteoDAO) GetSorteoByID(id uint) (*domain.Sorteo, error) {
	if m.getByIDErr != nil {
		return nil, m.getByIDErr
	}
	return m.sorteo, nil
}
func (m *mockSorteoDAO) GetSorteoByEventID(eventID uint) (*domain.Sorteo, error) {
	if m.getByEvtErr != nil {
		return nil, m.getByEvtErr
	}
	return m.sorteo, nil
}
func (m *mockSorteoDAO) GetActiveSorteoByEventID(eventID uint) (*domain.Sorteo, error) {
	if m.getActiveErr != nil {
		return nil, m.getActiveErr
	}
	if m.activeSorteo == nil {
		return nil, assert.AnError
	}
	return m.activeSorteo, nil
}
func (m *mockSorteoDAO) GetSorteosByEventID(eventID uint) ([]domain.Sorteo, error) {
	return m.sorteosByEvent, nil
}
func (m *mockSorteoDAO) UpdateSorteo(s *domain.Sorteo) error {
	m.lastSaved = s
	return m.updateErr
}
func (m *mockSorteoDAO) GetSorteosConEvento() ([]domain.Sorteo, error) {
	return m.sorteos, nil
}

type mockChanceDAO struct {
	chances    []domain.Chance
	createErr  error
	created    int
	countValue int64
}

func (m *mockChanceDAO) CreateChance(c *domain.Chance) error {
	m.created++
	c.IDChance = uint(m.created)
	return m.createErr
}
func (m *mockChanceDAO) GetChancesBySorteoID(sorteoID uint) ([]domain.Chance, error) {
	return m.chances, nil
}
func (m *mockChanceDAO) CountChancesByUserAndSorteo(userID, sorteoID uint) (int64, error) {
	return m.countValue, nil
}

type mockSorteoEventDAO struct {
	event  *domain.Event
	getErr error
}

func (m *mockSorteoEventDAO) GetEventByID(id uint) (*domain.Event, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.event, nil
}

type mockSorteoTicketDAO struct {
	ticket *domain.Ticket
	getErr error
}

func (m *mockSorteoTicketDAO) GetActiveTicketByUserAndEvent(userID, eventID uint) (*domain.Ticket, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	return m.ticket, nil
}

type mockSorteoUserDAO struct {
	users map[uint]*domain.User
}

func (m *mockSorteoUserDAO) GetUserByID(id uint) (*domain.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, assert.AnError
	}
	return u, nil
}

type mockSorteoNotificationDAO struct {
	created []domain.Notification
}

func (m *mockSorteoNotificationDAO) CreateNotification(n *domain.Notification) error {
	m.created = append(m.created, *n)
	return nil
}

// --- Tests: CreateSorteo ---

func TestCreateSorteo_Success(t *testing.T) {
	eventDAO := &mockSorteoEventDAO{event: &domain.Event{IDEvents: 1}}
	sorteoDAO := &mockSorteoDAO{getByEvtErr: assert.AnError}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, eventDAO, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	sorteo, err := svc.CreateSorteo(1, "Rifa solidaria", 500)
	assert.NoError(t, err)
	assert.Equal(t, "activo", sorteo.Estado)
	assert.Equal(t, "Rifa solidaria", sorteoDAO.sorteo.Nombre)
}

func TestCreateSorteo_EventNotFound(t *testing.T) {
	eventDAO := &mockSorteoEventDAO{getErr: assert.AnError}
	svc := services.NewSorteoService(&mockSorteoDAO{}, &mockChanceDAO{}, eventDAO, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.CreateSorteo(99, "Rifa", 500)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "evento no encontrado")
}

func TestCreateSorteo_ActiveAlreadyExists(t *testing.T) {
	eventDAO := &mockSorteoEventDAO{event: &domain.Event{IDEvents: 1}}
	sorteoDAO := &mockSorteoDAO{activeSorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Estado: "activo"}}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, eventDAO, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.CreateSorteo(1, "Rifa", 500)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ya tiene un sorteo activo")
}

func TestCreateSorteo_AllowedAfterPreviousRealizado(t *testing.T) {
	// Sin sorteo activo (el anterior ya fue realizado), se puede cargar uno nuevo para el mismo evento.
	eventDAO := &mockSorteoEventDAO{event: &domain.Event{IDEvents: 1}}
	sorteoDAO := &mockSorteoDAO{}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, eventDAO, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	sorteo, err := svc.CreateSorteo(1, "Segunda rifa", 300)
	assert.NoError(t, err)
	assert.Equal(t, "activo", sorteo.Estado)
}

func TestCreateSorteo_InvalidValor(t *testing.T) {
	svc := services.NewSorteoService(&mockSorteoDAO{}, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.CreateSorteo(1, "Rifa", 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "valor de la chance")
}

// --- Tests: BuyChances ---

func TestBuyChances_Success(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Estado: "activo"}}
	ticketDAO := &mockSorteoTicketDAO{ticket: &domain.Ticket{IDTickets: 1, IDUsers: 1, IDEvents: 1, Estado: "activo"}}
	chanceDAO := &mockChanceDAO{}
	notificationDAO := &mockSorteoNotificationDAO{}
	svc := services.NewSorteoService(sorteoDAO, chanceDAO, &mockSorteoEventDAO{}, ticketDAO, &mockSorteoUserDAO{}, notificationDAO)

	chances, err := svc.BuyChances(1, 1, 3)
	assert.NoError(t, err)
	assert.Len(t, chances, 3)
	assert.Equal(t, 3, chanceDAO.created)
	assert.Len(t, notificationDAO.created, 1, "debe crear una única notificación consolidada de la compra")
	assert.Equal(t, "chance_comprada", notificationDAO.created[0].Tipo)
}

func TestBuyChances_NoTicket(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Estado: "activo"}}
	ticketDAO := &mockSorteoTicketDAO{getErr: assert.AnError}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, ticketDAO, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.BuyChances(1, 1, 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "entrada activa")
}

func TestBuyChances_SorteoNotActive(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Estado: "realizado"}}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.BuyChances(1, 1, 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no admite nuevas chances")
}

func TestBuyChances_SorteoNotFound(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{getByIDErr: assert.AnError}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.BuyChances(1, 99, 1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sorteo no encontrado")
}

func TestBuyChances_InvalidCantidad(t *testing.T) {
	svc := services.NewSorteoService(&mockSorteoDAO{}, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.BuyChances(1, 1, 0)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "al menos 1")
}

// --- Tests: RunDraw ---

func TestRunDraw_Success(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Nombre: "Rifa", Estado: "activo"}}
	chanceDAO := &mockChanceDAO{chances: []domain.Chance{
		{IDChance: 1, IDSorteo: 1, IDUsers: 1},
		{IDChance: 2, IDSorteo: 1, IDUsers: 2},
	}}
	userDAO := &mockSorteoUserDAO{users: map[uint]*domain.User{
		1: {IDUsers: 1, Nombre: "Ana", Email: "ana@test.com"},
		2: {IDUsers: 2, Nombre: "Beto", Email: "beto@test.com"},
	}}
	notificationDAO := &mockSorteoNotificationDAO{}
	svc := services.NewSorteoService(sorteoDAO, chanceDAO, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, userDAO, notificationDAO)

	winner, err := svc.RunDraw(1)
	assert.NoError(t, err)
	assert.NotNil(t, winner)
	assert.Equal(t, "realizado", sorteoDAO.lastSaved.Estado)
	assert.NotNil(t, sorteoDAO.lastSaved.IDGanador)
	assert.Len(t, notificationDAO.created, 2, "debe notificar a ambos participantes (ganador y perdedor)")
	tipos := []string{notificationDAO.created[0].Tipo, notificationDAO.created[1].Tipo}
	assert.Contains(t, tipos, "sorteo_ganador")
	assert.Contains(t, tipos, "sorteo_perdedor")
}

func TestRunDraw_NoParticipants(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, Estado: "activo"}}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.RunDraw(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no tiene participantes")
}

func TestRunDraw_AlreadyRealizado(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, Estado: "realizado"}}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.RunDraw(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ya fue realizado")
}

func TestRunDraw_SorteoNotFound(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{getByIDErr: assert.AnError}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.RunDraw(99)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sorteo no encontrado")
}

// --- Tests: GetSorteoByEventID / GetMyChancesCount ---

func TestGetSorteoByEventID_NoSorteoCargado(t *testing.T) {
	// No tener sorteo es un estado válido (la mayoría de los eventos no tienen uno):
	// no debe propagarse como error.
	sorteoDAO := &mockSorteoDAO{getByEvtErr: gorm.ErrRecordNotFound}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	sorteo, err := svc.GetSorteoByEventID(1)
	assert.NoError(t, err)
	assert.Nil(t, sorteo)
}

func TestGetSorteoByEventID_DBError(t *testing.T) {
	// Un fallo real de la base de datos sí debe propagarse como error.
	sorteoDAO := &mockSorteoDAO{getByEvtErr: assert.AnError}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	_, err := svc.GetSorteoByEventID(1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error al obtener el sorteo")
}

func TestGetMyChancesCount(t *testing.T) {
	chanceDAO := &mockChanceDAO{countValue: 5}
	svc := services.NewSorteoService(&mockSorteoDAO{}, chanceDAO, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	count, err := svc.GetMyChancesCount(1, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), count)
}

// --- Tests: GetSorteosByEventID / GetChanceSummary ---

func TestGetSorteosByEventID(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteosByEvent: []domain.Sorteo{
		{IDSorteo: 2, IDEvents: 1, Estado: "activo"},
		{IDSorteo: 1, IDEvents: 1, Estado: "realizado"},
	}}
	svc := services.NewSorteoService(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	sorteos, err := svc.GetSorteosByEventID(1)
	assert.NoError(t, err)
	assert.Len(t, sorteos, 2)
}

func TestGetChanceSummary_GroupsByUser(t *testing.T) {
	chanceDAO := &mockChanceDAO{chances: []domain.Chance{
		{IDChance: 1, IDSorteo: 1, IDUsers: 1, User: domain.User{IDUsers: 1, Nombre: "Ana", Email: "ana@test.com"}},
		{IDChance: 2, IDSorteo: 1, IDUsers: 1, User: domain.User{IDUsers: 1, Nombre: "Ana", Email: "ana@test.com"}},
		{IDChance: 3, IDSorteo: 1, IDUsers: 2, User: domain.User{IDUsers: 2, Nombre: "Beto", Email: "beto@test.com"}},
	}}
	svc := services.NewSorteoService(&mockSorteoDAO{}, chanceDAO, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{}, &mockSorteoNotificationDAO{})

	summary, err := svc.GetChanceSummary(1)
	assert.NoError(t, err)
	assert.Len(t, summary, 2, "debe agrupar por usuario, no una fila por chance")
	assert.Equal(t, "Ana", summary[0].Nombre)
	assert.Equal(t, 2, summary[0].Cantidad)
	assert.Equal(t, "Beto", summary[1].Nombre)
	assert.Equal(t, 1, summary[1].Cantidad)
}
