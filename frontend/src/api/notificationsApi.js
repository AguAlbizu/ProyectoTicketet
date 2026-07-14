import api from './axiosConfig'

export const getMyNotifications = () =>
  api.get('/notifications')

export const markNotificationAsRead = (notificationId) =>
  api.put(`/notifications/${notificationId}/read`)

export const markAllNotificationsAsRead = () =>
  api.put('/notifications/read-all')
