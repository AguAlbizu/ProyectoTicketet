import { Link } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'

// Top navigation bar shown on all pages.
// Displays different links based on authentication state and role.
function Navbar() {
  const { user, logout } = useAuth()

  // TODO: render logo/brand link to "/"
  // TODO: if !user: show "Iniciar sesión" and "Registrarse" links
  // TODO: if user.role === "admin": show admin dashboard link
  // TODO: if user.role === "cliente": show "Mis entradas" link
  // TODO: show "Cerrar sesión" button that calls logout()

  return <nav>{/* TODO: implement Navbar UI */}</nav>
}

export default Navbar
