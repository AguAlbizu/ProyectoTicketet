import { Navigate, Outlet } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'
import LoadingSpinner from './LoadingSpinner'

// Redirige a /login si no hay token.
// Si se pasa requiredRole, redirige a / si el usuario no tiene ese rol.
function ProtectedRoute({ requiredRole }) {
  const { user, token, loading } = useAuth()

  if (loading) return <LoadingSpinner />
  if (!token) return <Navigate to="/login" replace />
  if (requiredRole && user?.rol !== requiredRole) return <Navigate to="/" replace />

  return <Outlet />
}

export default ProtectedRoute
