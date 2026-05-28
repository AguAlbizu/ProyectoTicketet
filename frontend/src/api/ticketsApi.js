import api from './axiosConfig'

// POST /tickets — purchases a ticket for an event
// payload: { event_id }
export const purchaseTicket = (payload) => api.post('/tickets', payload)

// GET /tickets/my — returns the authenticated user's tickets
export const getMyTickets = () => api.get('/tickets/my')

// DELETE /tickets/:id — cancels a ticket
export const cancelTicket = (id) => api.delete(`/tickets/${id}`)

// POST /tickets/:id/transfer — transfers a ticket to another user
// payload: { target_email }
export const transferTicket = (id, payload) =>
  api.post(`/tickets/${id}/transfer`, payload)
