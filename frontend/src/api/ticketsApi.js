import api from './axiosConfig'

export const buyTicket = (eventId, force = false) =>
  api.post('/tickets', { event_id: eventId, force })

export const getMyTickets = () =>
  api.get('/tickets/my-tickets')

export const cancelTicket = (ticketId) =>
  api.delete(`/tickets/${ticketId}`)

export const transferTicket = (ticketId, targetEmail) =>
  api.put(`/tickets/${ticketId}/transfer`, { target_email: targetEmail })
