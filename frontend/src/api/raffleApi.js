import api from './axiosConfig'

// GET /raffles/event/:eventId — returns the raffle for a given event
export const getRaffleByEvent = (eventId) =>
  api.get(`/raffles/event/${eventId}`)

// POST /raffles — creates a new raffle (admin only)
// payload: { event_id, name, price_per_chance }
export const createRaffle = (payload) => api.post('/raffles', payload)

// POST /raffles/:id/chances — buys chances for the authenticated user
// payload: { quantity }
export const buyChances = (raffleId, payload) =>
  api.post(`/raffles/${raffleId}/chances`, payload)

// POST /raffles/:id/draw — executes the raffle draw (admin only)
export const drawRaffle = (raffleId) =>
  api.post(`/raffles/${raffleId}/draw`)
