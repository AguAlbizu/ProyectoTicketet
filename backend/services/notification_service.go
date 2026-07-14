package services

import "ticketapp/domain"

// NotificationDAOPort define los métodos de persistencia de notificaciones.
type NotificationDAOPort interface {
	CreateNotification(n *domain.Notification) error
	GetByUserID(userID uint) ([]domain.Notification, error)
	MarkAsRead(id, userID uint) error
	MarkAllAsRead(userID uint) error
}

type NotificationService struct {
	notificationDAO NotificationDAOPort
}

func NewNotificationService(notificationDAO NotificationDAOPort) *NotificationService {
	return &NotificationService{notificationDAO: notificationDAO}
}

// GetMyNotifications retorna las notificaciones del usuario, más recientes primero.
func (s *NotificationService) GetMyNotifications(userID uint) ([]domain.Notification, error) {
	return s.notificationDAO.GetByUserID(userID)
}

// MarkAsRead marca una notificación puntual como leída (solo si le pertenece al usuario).
func (s *NotificationService) MarkAsRead(id, userID uint) error {
	return s.notificationDAO.MarkAsRead(id, userID)
}

// MarkAllAsRead marca todas las notificaciones no leídas del usuario como leídas.
func (s *NotificationService) MarkAllAsRead(userID uint) error {
	return s.notificationDAO.MarkAllAsRead(userID)
}
