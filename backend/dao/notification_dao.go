package dao

import (
	"ticketapp/domain"

	"gorm.io/gorm"
)

type NotificationDAO struct {
	db *gorm.DB
}

func NewNotificationDAO(db *gorm.DB) *NotificationDAO {
	return &NotificationDAO{db: db}
}

func (d *NotificationDAO) CreateNotification(n *domain.Notification) error {
	return d.db.Create(n).Error
}

// GetByUserID retorna las notificaciones del usuario, más recientes primero.
func (d *NotificationDAO) GetByUserID(userID uint) ([]domain.Notification, error) {
	var notifications []domain.Notification
	err := d.db.Where("id_users = ?", userID).Order("created_at DESC").Find(&notifications).Error
	return notifications, err
}

func (d *NotificationDAO) MarkAsRead(id, userID uint) error {
	return d.db.Exec(
		"UPDATE notifications SET leida = true WHERE id_notification = ? AND id_users = ?",
		id, userID,
	).Error
}

func (d *NotificationDAO) MarkAllAsRead(userID uint) error {
	return d.db.Exec(
		"UPDATE notifications SET leida = true WHERE id_users = ? AND leida = false",
		userID,
	).Error
}
