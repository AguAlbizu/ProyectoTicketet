import api from './axiosConfig'

export const buyTicket = (eventId) =>
  api.post('/tickets', { event_id: eventId })

export const getMyTickets = () =>
  api.get('/tickets/my-tickets')

export const cancelTicket = (ticketId) =>
  api.delete(`/tickets/${ticketId}`)

export const transferTicket = (ticketId, targetEmail) =>
  api.put(`/tickets/${ticketId}/transfer`, { target_email: targetEmail })
