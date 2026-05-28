import api from './axiosConfig'

// POST /auth/register — registers a new user
// payload: { name, email, password, role? }
export const register = (payload) => api.post('/auth/register', payload)

// POST /auth/login — authenticates a user and returns { token, user }
// payload: { email, password }
export const login = (payload) => api.post('/auth/login', payload)
