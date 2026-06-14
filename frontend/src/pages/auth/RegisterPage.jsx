import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { useAuth } from '../../hooks/useAuth'

function RegisterPage() {
  const { register } = useAuth()
  const navigate = useNavigate()
  const [nombre, setNombre] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      await register(nombre, email, password)
      navigate('/')
    } catch (err) {
      setError(err.response?.data?.error || 'Error al registrar usuario')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="auth-wrapper">
      <div className="auth-card">
        <img src="/logo.png" alt="TicketApp" style={{ height: '48px', marginBottom: '1.25rem', display: 'block' }} />
        <h2 className="auth-title">Creá tu cuenta</h2>
        <p className="auth-subtitle">Empezá a comprar entradas en segundos</p>

        {error && <div className="alert alert-error">{error}</div>}

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label className="form-label" htmlFor="nombre">Nombre</label>
            <input
              className="form-input"
              id="nombre"
              type="text"
              value={nombre}
              onChange={(e) => setNombre(e.target.value)}
              required
              placeholder="Tu nombre"
            />
          </div>
          <div className="form-group">
            <label className="form-label" htmlFor="email">Email</label>
            <input
              className="form-input"
              id="email"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              autoComplete="email"
              placeholder="tu@email.com"
            />
          </div>
          <div className="form-group">
            <label className="form-label" htmlFor="password">Contraseña</label>
            <input
              className="form-input"
              id="password"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              placeholder="Mínimo 6 caracteres"
            />
          </div>
          <button type="submit" className="btn btn-primary btn-full" disabled={loading}>
            {loading ? 'Registrando...' : 'Crear cuenta'}
          </button>
        </form>

        <p className="auth-footer">
          ¿Ya tenés cuenta?{' '}
          <Link to="/login">Iniciá sesión</Link>
        </p>
      </div>
    </div>
  )
}

export default RegisterPage
