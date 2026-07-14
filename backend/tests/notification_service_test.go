package tests

// Tests unitarios de NotificationService (notificaciones in-app).
// Correr con: go test ./tests/... -coverpkg=ticketapp/services,ticketapp/utils,ticketapp/controllers -v

import (
	"testing"
	"ticketapp/domain"
	"ticketapp/services"

	"github.com/stretchr/testify/assert"
)

// --- Mocks ---

type mockNotificationDAO struct {
	notifications  []domain.Notification
	markReadCalls  []uint
	markAllCalls   []uint
	markAsReadErr  error
	markAllReadErr error
}

func (m *mockNotificationDAO) CreateNotification(n *domain.Notification) error {
	m.notifications = append(m.notifications, *n)
	return nil
}

func (m *mockNotificationDAO) GetByUserID(userID uint) ([]domain.Notification, error) {
	var result []domain.Notification
	for _, n := range m.notifications {
		if n.IDUsers == userID {
			result = append(result, n)
		}
	}
	return result, nil
}

func (m *mockNotificationDAO) MarkAsRead(id, userID uint) error {
	m.markReadCalls = append(m.markReadCalls, id)
	return m.markAsReadErr
}

func (m *mockNotificationDAO) MarkAllAsRead(userID uint) error {
	m.markAllCalls = append(m.markAllCalls, userID)
	return m.markAllReadErr
}

// --- Tests ---

func TestGetMyNotifications(t *testing.T) {
	dao := &mockNotificationDAO{notifications: []domain.Notification{
		{IDNotification: 1, IDUsers: 1, Tipo: "chance_comprada"},
		{IDNotification: 2, IDUsers: 2, Tipo: "sorteo_ganador"},
	}}
	svc := services.NewNotificationService(dao)

	notifications, err := svc.GetMyNotifications(1)
	assert.NoError(t, err)
	assert.Len(t, notifications, 1)
	assert.Equal(t, "chance_comprada", notifications[0].Tipo)
}

func TestMarkAsRead(t *testing.T) {
	dao := &mockNotificationDAO{}
	svc := services.NewNotificationService(dao)

	err := svc.MarkAsRead(5, 1)
	assert.NoError(t, err)
	assert.Equal(t, []uint{5}, dao.markReadCalls)
}

func TestMarkAllAsRead(t *testing.T) {
	dao := &mockNotificationDAO{}
	svc := services.NewNotificationService(dao)

	err := svc.MarkAllAsRead(1)
	assert.NoError(t, err)
	assert.Equal(t, []uint{1}, dao.markAllCalls)
}
