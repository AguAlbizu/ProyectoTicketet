import { Navigate } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'

// Wraps a route to require authentication and optionally a specific role.
// Redirects to /login if unauthenticated, or / if role does not match.
//
// Props:
//   children     — the protected page component to render
//   requiredRole — optional: "cliente" | "admin"
function ProtectedRoute({ children, requiredRole }) {
  const { user } = useAuth()

  // TODO: if !user, return <Navigate to="/login" replace />
  // TODO: if requiredRole && user.role !== requiredRole, return <Navigate to="/" replace />
  // TODO: return children

  if (!user) return <Navigate to="/login" replace />
  if (requiredRole && user.role !== requiredRole) return <Navigate to="/" replace />
  return children
}

export default ProtectedRoute
