import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getEventById } from '../../api/eventsApi'
import { buyTicket } from '../../api/ticketsApi'
import Navbar from '../../components/common/Navbar'
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
      await buyTicket(id)
      navigate('/purchase-success')
    } catch (err) {
      setBuyError(err.response?.data?.error || 'Error al comprar la entrada')
    } finally {
      setBuying(false)
    }
  }

  if (loading) return <><Navbar /><LoadingSpinner /></>
  if (error) return <><Navbar /><p style={{ padding: '2rem', color: 'red' }}>{error}</p></>

  const fecha = event.fecha ? new Date(event.fecha).toLocaleDateString('es-AR') : ''

  return (
    <div>
      <Navbar />
      <main style={{ padding: '1rem 2rem', maxWidth: '800px', margin: '0 auto' }}>
        {event.imagen_url && (
          <img src={event.imagen_url} alt={event.titulo} style={{ width: '100%', maxHeight: '350px', objectFit: 'cover', borderRadius: '8px' }} />
        )}
        <h1>{event.titulo}</h1>
        <p><strong>Fecha:</strong> {fecha} — <strong>Hora:</strong> {event.hora}</p>
        <p><strong>Categoría:</strong> {event.categoria}</p>
        <p><strong>Capacidad:</strong> {event.capacidad} — <strong>Cupo disponible:</strong> {event.cupo_disponible}</p>
        <p>{event.descripcion}</p>

        {buyError && <p style={{ color: 'red' }}>{buyError}</p>}

        {event.cupo_disponible === 0 ? (
          <button disabled style={{ padding: '0.75rem 2rem', marginTop: '1rem' }}>Sin cupo</button>
        ) : (
          <button onClick={handlePurchase} disabled={buying} style={{ padding: '0.75rem 2rem', marginTop: '1rem' }}>
            {buying ? 'Procesando...' : 'Comprar Entrada'}
          </button>
        )}
      </main>
    </div>
  )
}

export default EventDetailPage
