import api from './axiosConfig'

// GET /events?category= — returns list of active events
export const getEvents = (category = '') =>
  api.get('/events', { params: category ? { category } : {} })

// GET /events/:id — returns a single event by ID
export const getEventById = (id) => api.get(`/events/${id}`)

// POST /events — creates a new event (admin only)
// payload: { title, description, date, duration_minutes, capacity, category, image_url }
export const createEvent = (payload) => api.post('/events', payload)

// PUT /events/:id — updates an existing event (admin only)
export const updateEvent = (id, payload) => api.put(`/events/${id}`, payload)

// DELETE /events/:id — cancels an event (admin only)
export const cancelEvent = (id) => api.delete(`/events/${id}`)
