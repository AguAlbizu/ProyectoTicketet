import { Navigate } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'

// Protege una ruta verificando que el usuario tenga un token válido en el contexto.
// Si no hay token, redirige a /login.
//
// TODO (entrega final): agregar validación de rol como prop (ej: requiredRole="admin")

function ProtectedRoute({ children }) {
  const { token } = useAuth()

  if (!token) {
    return <Navigate to="/login" replace />
  }

  return children
}

export default ProtectedRoute
