package tests

// Tests de integración HTTP para los endpoints de notificaciones.
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
)

func setupNotificationTestRouter(dao *mockNotificationDAO) *gin.Engine {
	svc := services.NewNotificationService(dao)
	ctrl := controllers.NewNotificationController(svc)
	r := gin.New()

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/notifications", ctrl.GetMyNotifications)
	protected.PUT("/notifications/:id/read", ctrl.MarkAsRead)
	protected.PUT("/notifications/read-all", ctrl.MarkAllAsRead)

	return r
}

func TestGetMyNotificationsEndpoint_NoToken(t *testing.T) {
	r := setupNotificationTestRouter(&mockNotificationDAO{})

	req := httptest.NewRequest(http.MethodGet, "/api/notifications", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetMyNotificationsEndpoint_Success(t *testing.T) {
	dao := &mockNotificationDAO{notifications: []domain.Notification{
		{IDNotification: 1, IDUsers: 1, Tipo: "chance_comprada"},
	}}
	r := setupNotificationTestRouter(dao)

	req := httptest.NewRequest(http.MethodGet, "/api/notifications", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMarkAsReadEndpoint_Success(t *testing.T) {
	dao := &mockNotificationDAO{}
	r := setupNotificationTestRouter(dao)

	req := httptest.NewRequest(http.MethodPut, "/api/notifications/1/read", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, []uint{1}, dao.markReadCalls)
}

func TestMarkAllAsReadEndpoint_Success(t *testing.T) {
	dao := &mockNotificationDAO{}
	r := setupNotificationTestRouter(dao)

	req := httptest.NewRequest(http.MethodPut, "/api/notifications/read-all", nil)
	req.Header.Set("Authorization", "Bearer "+testToken(t))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Len(t, dao.markAllCalls, 1)
}
