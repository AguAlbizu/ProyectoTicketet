import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getEventById } from '../../api/eventsApi'
import { buyTicket } from '../../api/ticketsApi'
import Navbar from '../../components/common/Navbar'
import Footer from '../../components/common/Footer'
import LoadingSpinner from '../../components/common/LoadingSpinner'
import { useAuth } from '../../hooks/useAuth'

function EventDetailPage() {
  const { id } = useParams()
  const { isAuthenticated } = useAuth()
  const navigate = useNavigate()

  const [event, setEvent] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [buying, setBuying] = useState(false)
  const [buyError, setBuyError] = useState('')

  useEffect(() => {
    const fetchEvent = async () => {
      try {
        const res = await getEventById(id)
        setEvent(res.data)
      } catch {
        setError('Evento no encontrado')
      } finally {
        setLoading(false)
      }
    }
    fetchEvent()
  }, [id])

  const handlePurchase = async () => {
    if (!isAuthenticated) {
      navigate('/login')
      return
    }
    setBuyError('')
    setBuying(true)
    try {
      await buyTicket(Number(id))
      navigate('/purchase-success')
    } catch (err) {
      setBuyError(err.response?.data?.error || 'Error al comprar la entrada')
    } finally {
      setBuying(false)
    }
  }

  if (loading) return <div className="page-wrapper"><Navbar /><LoadingSpinner /></div>

  if (error) return (
    <div className="page-wrapper">
      <Navbar />
      <main className="page-main">
        <div className="alert alert-error">{error}</div>
        <button className="btn btn-outline" onClick={() => navigate('/')}>Volver al catálogo</button>
      </main>
    </div>
  )

  const fecha = event.fecha
    ? new Date(event.fecha).toLocaleDateString('es-AR', { weekday: 'long', day: 'numeric', month: 'long', year: 'numeric' })
    : ''

  return (
    <div className="page-wrapper">
      <Navbar />
      <main className="page-main-narrow">
        {event.imagen_url && (
          <img className="event-detail-img" src={event.imagen_url} alt={event.titulo} />
        )}

        <span className="badge badge-category" style={{ marginBottom: '0.75rem', display: 'inline-block' }}>
          {event.categoria}
        </span>
        <h1 className="page-title">{event.titulo}</h1>

        <div className="event-detail-meta">
          {fecha && (
            <div className="event-detail-meta-item">
              <span className="event-detail-meta-label">Fecha</span>
              <span className="event-detail-meta-value">{fecha}</span>
            </div>
          )}
          <div className="event-detail-meta-item">
            <span className="event-detail-meta-label">Hora</span>
            <span className="event-detail-meta-value">{event.hora}</span>
          </div>
          <div className="event-detail-meta-item">
            <span className="event-detail-meta-label">Cupo disponible</span>
            <span className="event-detail-meta-value">{event.cupo_disponible} de {event.capacidad}</span>
          </div>
          {event.direccion && (
            <div className="event-detail-meta-item">
              <span className="event-detail-meta-label">Lugar</span>
              <span className="event-detail-meta-value">{event.direccion}</span>
            </div>
          )}
        </div>

        {event.descripcion && (
          <p className="event-detail-desc">{event.descripcion}</p>
        )}

        <div className="event-detail-buy">
          {buyError && <div className="alert alert-error">{buyError}</div>}

          <div className="event-detail-buy-row">
            {event.precio > 0 && (
              <div className="event-detail-price-wrapper">
                <p className="event-detail-price">
                  ${event.precio.toLocaleString('es-AR')}
                </p>
                <span className="event-detail-price-label">por entrada</span>
              </div>
            )}

            {event.cupo_disponible === 0 ? (
              <span className="event-detail-soldout">Sin cupo disponible</span>
            ) : (
              <button
                className="btn btn-primary btn-lg"
                onClick={handlePurchase}
                disabled={buying}
              >
                {buying ? 'Procesando...' : 'Comprar entrada'}
              </button>
            )}
          </div>
        </div>
      </main>
      <Footer />
    </div>
  )
}

export default EventDetailPage
