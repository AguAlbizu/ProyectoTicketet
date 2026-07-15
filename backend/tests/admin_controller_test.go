package tests

// Tests de integración HTTP para los endpoints de administrador (eventos y usuarios).
// Correr con: go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils,ticketapp/controllers -v

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"ticketapp/controllers"
	"ticketapp/domain"
	"ticketapp/middleware"
	"ticketapp/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupAdminTestRouter(eventDAO *mockAdminEventDAO, ticketDAO *mockAdminTicketDAO, authDAO *mockAuthUserDAO) *gin.Engine {
	adminService := services.NewAdminEventService(eventDAO, ticketDAO)
	authService := services.NewAuthService(authDAO)
	ctrl := controllers.NewAdminController(adminService, authService)
	r := gin.New()

	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RequireRole("administrador"))
	admin.GET("/events", ctrl.GetAllEvents)
	admin.POST("/events", ctrl.CreateEvent)
	admin.PUT("/events/:id", ctrl.UpdateEvent)
	admin.DELETE("/events/:id", ctrl.CancelEvent)
	admin.GET("/events/:id/report", ctrl.GetEventReport)
	admin.POST("/users", ctrl.CreateAdmin)
	admin.PUT("/users/promote", ctrl.PromoteToAdmin)

	return r
}

// =============================================
// Autorización (compartida por todos los endpoints admin)
// =============================================

func TestAdminEndpoint_ForbiddenForClient(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/admin/events", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t)) // token de cliente
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

// =============================================
// PromoteToAdmin
// =============================================

func TestPromoteToAdminEndpoint_Success(t *testing.T) {
	authDAO := &mockAuthUserDAO{user: &domain.User{Email: "juan@test.com", Rol: "cliente"}}
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, authDAO)

	req := httptest.NewRequest(http.MethodPut, "/api/admin/users/promote", jsonBody(`{"email":"juan@test.com"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPromoteToAdminEndpoint_MissingEmail(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodPut, "/api/admin/users/promote", jsonBody(`{}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestPromoteToAdminEndpoint_NotFound(t *testing.T) {
	authDAO := &mockAuthUserDAO{findErr: assert.AnError}
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, authDAO)

	req := httptest.NewRequest(http.MethodPut, "/api/admin/users/promote", jsonBody(`{"email":"noexiste@test.com"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPromoteToAdminEndpoint_AlreadyAdmin(t *testing.T) {
	authDAO := &mockAuthUserDAO{user: &domain.User{Email: "admin@test.com", Rol: "administrador"}}
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, authDAO)

	req := httptest.NewRequest(http.MethodPut, "/api/admin/users/promote", jsonBody(`{"email":"admin@test.com"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

// =============================================
// CreateAdmin
// =============================================

func TestCreateAdminEndpoint_Success(t *testing.T) {
	authDAO := &mockAuthUserDAO{findErr: gorm.ErrRecordNotFound}
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, authDAO)

	req := httptest.NewRequest(http.MethodPost, "/api/admin/users",
		jsonBody(`{"nombre":"Root","email":"root@test.com","password":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateAdminEndpoint_MissingFields(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/admin/users", jsonBody(`{"email":"root@test.com"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateAdminEndpoint_DuplicateEmail(t *testing.T) {
	authDAO := &mockAuthUserDAO{user: &domain.User{Email: "root@test.com"}}
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, authDAO)

	req := httptest.NewRequest(http.MethodPost, "/api/admin/users",
		jsonBody(`{"nombre":"Root","email":"root@test.com","password":"123456"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
}

// =============================================
// GetAllEvents (admin)
// =============================================

func TestAdminGetAllEventsEndpoint_Success(t *testing.T) {
	eventDAO := &mockAdminEventDAO{events: []domain.Event{{IDEvents: 1, Titulo: "Evento A"}}}
	r := setupAdminTestRouter(eventDAO, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/admin/events", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// =============================================
// CreateEvent (admin)
// =============================================

func TestAdminCreateEventEndpoint_Success(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	body := `{"titulo":"Recital","fecha":"2026-12-01T00:00:00Z","hora":"20:00","capacidad":100}`
	req := httptest.NewRequest(http.MethodPost, "/api/admin/events", jsonBody(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestAdminCreateEventEndpoint_MissingFields(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodPost, "/api/admin/events", jsonBody(`{"titulo":"Sin fecha"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdminCreateEventEndpoint_DAOError(t *testing.T) {
	eventDAO := &mockAdminEventDAO{createErr: assert.AnError}
	r := setupAdminTestRouter(eventDAO, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	body := `{"titulo":"Recital","fecha":"2026-12-01T00:00:00Z","hora":"20:00","capacidad":100}`
	req := httptest.NewRequest(http.MethodPost, "/api/admin/events", jsonBody(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

// =============================================
// UpdateEvent (admin)
// =============================================

func TestAdminUpdateEventEndpoint_Success(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo", Capacidad: 100, CupoDisponible: 100}}
	r := setupAdminTestRouter(eventDAO, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	body := `{"titulo":"Recital actualizado","fecha":"2026-12-01T00:00:00Z","hora":"21:00","capacidad":150}`
	req := httptest.NewRequest(http.MethodPut, "/api/admin/events/1", jsonBody(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAdminUpdateEventEndpoint_InvalidID(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodPut, "/api/admin/events/abc", jsonBody(`{}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdminUpdateEventEndpoint_NotFound(t *testing.T) {
	eventDAO := &mockAdminEventDAO{getErr: assert.AnError}
	r := setupAdminTestRouter(eventDAO, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	body := `{"titulo":"X","fecha":"2026-12-01T00:00:00Z","hora":"21:00","capacidad":100}`
	req := httptest.NewRequest(http.MethodPut, "/api/admin/events/99", jsonBody(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// =============================================
// CancelEvent (admin)
// =============================================

func TestAdminCancelEventEndpoint_Success(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Estado: "activo"}}
	r := setupAdminTestRouter(eventDAO, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodDelete, "/api/admin/events/1", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAdminCancelEventEndpoint_InvalidID(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodDelete, "/api/admin/events/abc", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdminCancelEventEndpoint_NotFound(t *testing.T) {
	eventDAO := &mockAdminEventDAO{getErr: assert.AnError}
	r := setupAdminTestRouter(eventDAO, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodDelete, "/api/admin/events/99", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// =============================================
// GetEventReport (admin)
// =============================================

func TestAdminGetEventReportEndpoint_Success(t *testing.T) {
	eventDAO := &mockAdminEventDAO{event: &domain.Event{IDEvents: 1, Titulo: "Evento", Capacidad: 100, CupoDisponible: 90}}
	ticketDAO := &mockAdminTicketDAO{tickets: []domain.Ticket{
		{IDTickets: 1, IDUsers: 1, Estado: "activo", Origen: "compra"},
	}}
	r := setupAdminTestRouter(eventDAO, ticketDAO, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/admin/events/1/report", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAdminGetEventReportEndpoint_InvalidID(t *testing.T) {
	r := setupAdminTestRouter(&mockAdminEventDAO{}, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/admin/events/abc/report", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAdminGetEventReportEndpoint_NotFound(t *testing.T) {
	eventDAO := &mockAdminEventDAO{getErr: assert.AnError}
	r := setupAdminTestRouter(eventDAO, &mockAdminTicketDAO{}, &mockAuthUserDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/admin/events/99/report", nil)
	req.Header.Set("Authorization", "Bearer "+adminToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
