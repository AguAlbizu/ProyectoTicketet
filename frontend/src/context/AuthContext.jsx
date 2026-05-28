import { createContext, useState } from 'react'
import * as authApi from '../api/authApi'

export const AuthContext = createContext(null)

export function AuthProvider({ children }) {
  const [user, setUser] = useState(() => {
    // Restore user from localStorage on page reload
    const stored = localStorage.getItem('user')
    return stored ? JSON.parse(stored) : null
  })

  const [token, setToken] = useState(() => localStorage.getItem('token') || null)

  // Register a new account and automatically log in on success.
  const register = async (name, email, password) => {
    // TODO: call authApi.register({ name, email, password })
    // TODO: on success, persist token and user to localStorage
    // TODO: update state: setToken, setUser
  }

  // Log in with email and password.
  // Persists the JWT token and user data to localStorage.
  const login = async (email, password) => {
    // TODO: call authApi.login({ email, password })
    // TODO: destructure response: { token, user }
    // TODO: localStorage.setItem('token', token), localStorage.setItem('user', JSON.stringify(user))
    // TODO: setToken(token), setUser(user)
  }

  // Clear session data and redirect to login.
  const logout = () => {
    // TODO: localStorage.removeItem('token'), localStorage.removeItem('user')
    // TODO: setToken(null), setUser(null)
  }

  return (
    <AuthContext.Provider value={{ user, token, login, logout, register }}>
      {children}
    </AuthContext.Provider>
  )
}
