import { useContext } from 'react'
import { AuthContext } from '../context/AuthContext'

// Convenience hook for consuming AuthContext.
// Returns { user, token, login, logout, register }.
export function useAuth() {
  const context = useContext(AuthContext)
  if (!context) {
    throw new Error('useAuth must be used inside an AuthProvider')
  }
  return context
}
