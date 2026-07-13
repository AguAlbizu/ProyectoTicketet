package tests

// Tests de integración HTTP para los endpoints de sorteos (Bonus Track).
// Correr con: go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils,ticketapp/controllers -v

import (
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
)

func setupSorteoTestRouter(sorteoDAO *mockSorteoDAO, chanceDAO *mockChanceDAO, eventDAO *mockSorteoEventDAO, ticketDAO *mockSorteoTicketDAO, userDAO *mockSorteoUserDAO) *gin.Engine {
	svc := services.NewSorteoService(sorteoDAO, chanceDAO, eventDAO, ticketDAO, userDAO, &mockEmailClient{})
	ctrl := controllers.NewSorteoController(svc)
	r := gin.New()

	r.GET("/api/events/:id/sorteo", ctrl.GetSorteoByEvent)

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/sorteos/:id/chances", ctrl.BuyChances)
	protected.GET("/sorteos/:id/my-chances", ctrl.GetMyChances)

	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("administrador"))
	admin.POST("/events/:id/sorteo", ctrl.CreateSorteo)
	admin.GET("/sorteos", ctrl.ListSorteosAdmin)
	admin.POST("/sorteos/:id/draw", ctrl.RunDraw)

	return r
}

func adminToken(t *testing.T) string {
	os.Setenv("JWT_SECRET", "test-secret-controllers")
	token, err := utils.GenerateToken(1, "administrador", "admin@test.com")
	assert.NoError(t, err)
	return token
}

// =============================================
// SORTEO CONTROLLER — endpoints de cliente
// =============================================

func TestGetSorteoByEventEndpoint_Success(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Nombre: "Rifa", Estado: "activo"}}
	r := setupSorteoTestRouter(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/events/1/sorteo", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetSorteoByEventEndpoint_NotFound(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{getByEvtErr: assert.AnError}
	r := setupSorteoTestRouter(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/events/1/sorteo", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBuyChancesEndpoint_NoToken(t *testing.T) {
	r := setupSorteoTestRouter(&mockSorteoDAO{}, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/sorteos/1/chances", jsonBody(`{"cantidad":1}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestBuyChancesEndpoint_Success(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Estado: "activo"}}
	ticketDAO := &mockSorteoTicketDAO{ticket: &domain.Ticket{IDTickets: 1, IDUsers: 1, IDEvents: 1, Estado: "activo"}}
	r := setupSorteoTestRouter(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, ticketDAO, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/sorteos/1/chances", jsonBody(`{"cantidad":2}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestBuyChancesEndpoint_NoActiveTicket(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Estado: "activo"}}
	ticketDAO := &mockSorteoTicketDAO{getErr: assert.AnError}
	r := setupSorteoTestRouter(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, ticketDAO, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/sorteos/1/chances", jsonBody(`{"cantidad":1}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetMyChancesEndpoint(t *testing.T) {
	chanceDAO := &mockChanceDAO{countValue: 4}
	r := setupSorteoTestRouter(&mockSorteoDAO{}, chanceDAO, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/sorteos/1/my-chances", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// =============================================
// SORTEO CONTROLLER — endpoints de administrador
// =============================================

func TestCreateSorteoEndpoint_Forbidden(t *testing.T) {
	eventDAO := &mockSorteoEventDAO{event: &domain.Event{IDEvents: 1}}
	r := setupSorteoTestRouter(&mockSorteoDAO{getByEvtErr: assert.AnError}, &mockChanceDAO{}, eventDAO, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/admin/events/1/sorteo", jsonBody(`{"nombre":"Rifa","valor_chance":500}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+testToken(t)) // token de cliente, no admin
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestCreateSorteoEndpoint_Success(t *testing.T) {
	eventDAO := &mockSorteoEventDAO{event: &domain.Event{IDEvents: 1}}
	r := setupSorteoTestRouter(&mockSorteoDAO{getByEvtErr: assert.AnError}, &mockChanceDAO{}, eventDAO, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/admin/events/1/sorteo", jsonBody(`{"nombre":"Rifa","valor_chance":500}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestRunDrawEndpoint_Forbidden(t *testing.T) {
	r := setupSorteoTestRouter(&mockSorteoDAO{}, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/admin/sorteos/1/draw", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t)) // token de cliente, no admin
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestRunDrawEndpoint_Success(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, IDEvents: 1, Nombre: "Rifa", Estado: "activo"}}
	chanceDAO := &mockChanceDAO{chances: []domain.Chance{{IDChance: 1, IDSorteo: 1, IDUsers: 1}}}
	userDAO := &mockSorteoUserDAO{users: map[uint]*domain.User{1: {IDUsers: 1, Nombre: "Ana", Email: "ana@test.com"}}}
	r := setupSorteoTestRouter(sorteoDAO, chanceDAO, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, userDAO)

	req := httptest.NewRequest(http.MethodPost, "/api/admin/sorteos/1/draw", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRunDrawEndpoint_NoParticipants(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteo: &domain.Sorteo{IDSorteo: 1, Estado: "activo"}}
	r := setupSorteoTestRouter(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/admin/sorteos/1/draw", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestListSorteosAdminEndpoint(t *testing.T) {
	sorteoDAO := &mockSorteoDAO{sorteos: []domain.Sorteo{
		{IDSorteo: 1, IDEvents: 1, Nombre: "Rifa A", Estado: "activo"},
		{IDSorteo: 2, IDEvents: 2, Nombre: "Rifa B", Estado: "realizado"},
	}}
	r := setupSorteoTestRouter(sorteoDAO, &mockChanceDAO{}, &mockSorteoEventDAO{}, &mockSorteoTicketDAO{}, &mockSorteoUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/admin/sorteos", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
