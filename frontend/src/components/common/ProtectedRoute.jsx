import { Navigate, Outlet } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'
import LoadingSpinner from './LoadingSpinner'

// Redirige a /login si no hay token.
// TODO (entrega final): agregar validación de rol como prop
function ProtectedRoute() {
  const { token, loading } = useAuth()

  if (loading) return <LoadingSpinner />
  if (!token) return <Navigate to="/login" replace />

  return <Outlet />
}

export default ProtectedRoute
