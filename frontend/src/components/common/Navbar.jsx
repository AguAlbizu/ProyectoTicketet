import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'

function Navbar() {
  const { user, isAuthenticated, logout } = useAuth()
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  return (
    <nav className="navbar">
      <Link to="/">
        <img src="/logo.png" alt="TicketApp" style={{ height: '58px', display: 'block' }} />
      </Link>

      <div className="navbar-links">
        {isAuthenticated ? (
          <>
            <span className="navbar-greeting">Hola, {user?.nombre}</span>
            <Link to="/my-tickets" className="navbar-link">Mis Entradas</Link>
            {user?.rol === 'administrador' && (
              <Link to="/admin/events" className="navbar-link">Panel Admin</Link>
            )}
            <button className="btn btn-ghost" onClick={handleLogout}>Salir</button>
          </>
        ) : (
          <>
            <Link to="/login" className="navbar-link">Iniciar Sesión</Link>
            <Link to="/register" className="btn btn-primary">Registrarse</Link>
          </>
        )}
      </div>
    </nav>
  )
}

export default Navbar
