import api from './axiosConfig'

export const getEvents = (categoria = '') =>
  api.get('/events', { params: categoria ? { categoria } : {} })

export const getEventById = (id) =>
  api.get(`/events/${id}`)
