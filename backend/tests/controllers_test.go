package tests

// Tests de integración HTTP usando net/http/httptest.
// Verifican que los endpoints retornan los códigos de estado correctos
// y que el middleware JWT funciona correctamente.
// Correr con: go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils,ticketapp/controllers -v

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"ticketapp/controllers"
	"ticketapp/domain"
	"ticketapp/middleware"
	"ticketapp/services"
	"ticketapp/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func init() {
	gin.SetMode(gin.TestMode)
	os.Setenv("JWT_SECRET", "test-secret-controllers")
	os.Setenv("JWT_EXPIRATION_HOURS", "24")
}

// --- Helpers ---

func setupAuthTestRouter(dao *mockAuthUserDAO) *gin.Engine {
	svc := services.NewAuthService(dao)
	ctrl := controllers.NewAuthController(svc)
	r := gin.New()
	r.POST("/api/auth/register", ctrl.Register)
	r.POST("/api/auth/login", ctrl.Login)
	return r
}

func setupEventTestRouter(dao services.EventRepository) *gin.Engine {
	svc := services.NewEventService(dao)
	ctrl := controllers.NewEventController(svc)
	r := gin.New()
	r.GET("/api/events", ctrl.GetEvents)
	r.GET("/api/events/:id", ctrl.GetEventByID)
	return r
}

func setupTicketTestRouter(ticketDAO *mockTicketDAO, eventDAO *mockEventDAOForTicket, userDAO *mockUserDAOForTicket) *gin.Engine {
	svc := services.NewTicketService(ticketDAO, eventDAO, userDAO)
	ctrl := controllers.NewTicketController(svc)
	r := gin.New()
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/tickets", ctrl.BuyTicket)
	protected.GET("/tickets/my-tickets", ctrl.GetMyTickets)
	protected.DELETE("/tickets/:id", ctrl.CancelTicket)
	return r
}

func testToken(t *testing.T) string {
	os.Setenv("JWT_SECRET", "test-secret-controllers")
	token, err := utils.GenerateToken(1, "cliente", "user@test.com")
	assert.NoError(t, err)
	return token
}

func jsonBody(s string) *bytes.Buffer {
	return bytes.NewBufferString(s)
}

// =============================================
// AUTH CONTROLLER
// =============================================

// TestRegisterEndpoint_Success verifica que POST /api/auth/register con datos válidos retorna 201.
func TestRegisterEndpoint_Success(t *testing.T) {
	dao := &mockAuthUserDAO{findErr: gorm.ErrRecordNotFound}
	r := setupAuthTestRouter(dao)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/register",
		jsonBody(`{"nombre":"Test","email":"test@test.com","password":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestRegisterEndpoint_MissingFields verifica que POST /api/auth/register sin campos requeridos retorna 400.
func TestRegisterEndpoint_MissingFields(t *testing.T) {
	r := setupAuthTestRouter(&mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/auth/register",
		jsonBody(`{"email":"test@test.com"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// TestRegisterEndpoint_DuplicateEmail verifica que registrar un email ya existente retorna 409.
func TestRegisterEndpoint_DuplicateEmail(t *testing.T) {
	dao := &mockAuthUserDAO{user: &domain.User{Email: "ya@existe.com"}}
	r := setupAuthTestRouter(dao)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/register",
		jsonBody(`{"nombre":"X","email":"ya@existe.com","password":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusConflict, w.Code)
}

// TestLoginEndpoint_Success verifica que POST /api/auth/login con credenciales válidas retorna 200 con token.
func TestLoginEndpoint_Success(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-controllers")
	hashed := utils.HashPassword("password123")
	dao := &mockAuthUserDAO{user: &domain.User{IDUsers: 1, Email: "u@test.com", Password: hashed, Rol: "cliente"}}
	r := setupAuthTestRouter(dao)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login",
		jsonBody(`{"email":"u@test.com","password":"password123"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var resp map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NotEmpty(t, resp["token"])
}

// TestLoginEndpoint_InvalidCredentials verifica que POST /api/auth/login con contraseña incorrecta retorna 401.
func TestLoginEndpoint_InvalidCredentials(t *testing.T) {
	hashed := utils.HashPassword("correcta")
	dao := &mockAuthUserDAO{user: &domain.User{IDUsers: 1, Email: "u@test.com", Password: hashed, Rol: "cliente"}}
	r := setupAuthTestRouter(dao)

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login",
		jsonBody(`{"email":"u@test.com","password":"incorrecta"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestLoginEndpoint_MissingFields verifica que POST /api/auth/login sin campos retorna 400.
func TestLoginEndpoint_MissingFields(t *testing.T) {
	r := setupAuthTestRouter(&mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/auth/login", jsonBody(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// =============================================
// EVENT CONTROLLER
// =============================================

// TestGetEventsEndpoint verifica que GET /api/events retorna 200 con la lista de eventos.
func TestGetEventsEndpoint(t *testing.T) {
	dao := &mockEventRepository{
		events: []domain.Event{
			{IDEvents: 1, Titulo: "Evento A", Estado: "activo", Categoria: "Música"},
			{IDEvents: 2, Titulo: "Evento B", Estado: "activo", Categoria: "Teatro"},
		},
	}
	r := setupEventTestRouter(dao)

	req := httptest.NewRequest(http.MethodGet, "/api/events", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var result []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Len(t, result, 2)
}

// TestGetEventsEndpoint_WithFilter verifica que GET /api/events?categoria=Música filtra correctamente.
func TestGetEventsEndpoint_WithFilter(t *testing.T) {
	dao := &mockEventRepository{
		events: []domain.Event{
			{IDEvents: 1, Titulo: "Concierto", Estado: "activo", Categoria: "Música"},
			{IDEvents: 2, Titulo: "Obra", Estado: "activo", Categoria: "Teatro"},
		},
	}
	r := setupEventTestRouter(dao)

	req := httptest.NewRequest(http.MethodGet, "/api/events?categoria=Música", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var result []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Len(t, result, 1)
}

// TestGetEventByIDEndpoint_Success verifica que GET /api/events/:id con ID válido retorna 200.
func TestGetEventByIDEndpoint_Success(t *testing.T) {
	dao := &mockEventRepository{
		events: []domain.Event{
			{IDEvents: 1, Titulo: "Evento Test", Estado: "activo"},
		},
	}
	r := setupEventTestRouter(dao)

	req := httptest.NewRequest(http.MethodGet, "/api/events/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestGetEventByIDEndpoint_NotFound verifica que GET /api/events/:id con ID inexistente retorna 404.
func TestGetEventByIDEndpoint_NotFound(t *testing.T) {
	dao := &mockEventRepository{findErr: gorm.ErrRecordNotFound}
	r := setupEventTestRouter(dao)

	req := httptest.NewRequest(http.MethodGet, "/api/events/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestGetEventByIDEndpoint_InvalidID verifica que GET /api/events/:id con ID no numérico retorna 400.
func TestGetEventByIDEndpoint_InvalidID(t *testing.T) {
	r := setupEventTestRouter(&mockEventRepository{})

	req := httptest.NewRequest(http.MethodGet, "/api/events/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// =============================================
// TICKET CONTROLLER
// =============================================

// TestBuyTicketEndpoint_NoToken verifica que POST /api/tickets sin token retorna 401.
func TestBuyTicketEndpoint_NoToken(t *testing.T) {
	r := setupTicketTestRouter(&mockTicketDAO{}, &mockEventDAOForTicket{}, &mockUserDAOForTicket{})

	req := httptest.NewRequest(http.MethodPost, "/api/tickets", jsonBody(`{"event_id":1}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestBuyTicketEndpoint_Success verifica que POST /api/tickets con token válido y cupo disponible retorna 201.
func TestBuyTicketEndpoint_Success(t *testing.T) {
	eventDAO := &mockEventDAOForTicket{
		event: &domain.Event{IDEvents: 1, Estado: "activo", CupoDisponible: 10},
	}
	r := setupTicketTestRouter(&mockTicketDAO{}, eventDAO, &mockUserDAOForTicket{})

	req := httptest.NewRequest(http.MethodPost, "/api/tickets", jsonBody(`{"event_id":1}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

// TestBuyTicketEndpoint_EventNotFound verifica que comprar en un evento inexistente retorna 404.
func TestBuyTicketEndpoint_EventNotFound(t *testing.T) {
	r := setupTicketTestRouter(&mockTicketDAO{}, &mockEventDAOForTicket{}, &mockUserDAOForTicket{})

	req := httptest.NewRequest(http.MethodPost, "/api/tickets", jsonBody(`{"event_id":999}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

// TestGetMyTicketsEndpoint verifica que GET /api/tickets/my-tickets con token retorna 200 con la lista.
func TestGetMyTicketsEndpoint(t *testing.T) {
	ticketDAO := &mockTicketDAO{
		tickets: []domain.Ticket{
			{IDTickets: 1, IDUsers: 1, Estado: "activo"},
			{IDTickets: 2, IDUsers: 1, Estado: "cancelado"},
		},
	}
	r := setupTicketTestRouter(ticketDAO, &mockEventDAOForTicket{event: &domain.Event{}}, &mockUserDAOForTicket{})

	req := httptest.NewRequest(http.MethodGet, "/api/tickets/my-tickets", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var result []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Len(t, result, 2)
}

// TestGetMyTicketsEndpoint_NoToken verifica que GET /api/tickets/my-tickets sin token retorna 401.
func TestGetMyTicketsEndpoint_NoToken(t *testing.T) {
	r := setupTicketTestRouter(&mockTicketDAO{}, &mockEventDAOForTicket{}, &mockUserDAOForTicket{})

	req := httptest.NewRequest(http.MethodGet, "/api/tickets/my-tickets", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestCancelTicketEndpoint_NotFound verifica que DELETE /api/tickets/:id con ticket inexistente retorna 404.
func TestCancelTicketEndpoint_NotFound(t *testing.T) {
	ticketDAO := &mockTicketDAO{getErr: gorm.ErrRecordNotFound}
	r := setupTicketTestRouter(ticketDAO, &mockEventDAOForTicket{}, &mockUserDAOForTicket{})

	req := httptest.NewRequest(http.MethodDelete, "/api/tickets/99", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
