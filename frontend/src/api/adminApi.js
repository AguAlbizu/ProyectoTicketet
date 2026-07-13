import api from './axiosConfig'

export const getAllEventsAdmin = () => api.get('/admin/events')
export const createEvent = (data) => api.post('/admin/events', data)
export const updateEvent = (id, data) => api.put(`/admin/events/${id}`, data)
export const cancelEvent = (id) => api.delete(`/admin/events/${id}`)
export const getEventReport = (id) => api.get(`/admin/events/${id}/report`)
export const createAdmin = (data) => api.post('/admin/users', data)
export const promoteToAdmin = (email) => api.put('/admin/users/promote', { email })
