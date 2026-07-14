import api from './axiosConfig'

// --- Cliente ---

export const getSorteoByEvent = (eventId) =>
  api.get(`/events/${eventId}/sorteo`)

export const buyChances = (sorteoId, cantidad) =>
  api.post(`/sorteos/${sorteoId}/chances`, { cantidad })

export const getMyChances = (sorteoId) =>
  api.get(`/sorteos/${sorteoId}/my-chances`)

// --- Administrador ---

export const createSorteo = (eventId, nombre, valorChance) =>
  api.post(`/admin/events/${eventId}/sorteo`, { nombre, valor_chance: valorChance })

export const listSorteosAdmin = () =>
  api.get('/admin/sorteos')

export const getSorteosByEvent = (eventId) =>
  api.get(`/admin/events/${eventId}/sorteos`)

export const runSorteoDraw = (sorteoId) =>
  api.post(`/admin/sorteos/${sorteoId}/draw`)

export const getSorteoChanceSummary = (sorteoId) =>
  api.get(`/admin/sorteos/${sorteoId}/chances`)
