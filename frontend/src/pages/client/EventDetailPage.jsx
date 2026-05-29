import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getEventById } from '../../api/eventsApi'
import { purchaseTicket } from '../../api/ticketsApi'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'
import { useAuth } from '../../hooks/useAuth'

// Página de detalle de un evento — muestra info completa y botón de compra de entrada.
function EventDetailPage() {
  const { id } = useParams()
  const { user } = useAuth()
  const navigate = useNavigate()

  const [event, setEvent] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    // TODO: fetch evento con getEventById(id)
    // TODO: setEvent(response.data), setLoading(false)
  }, [id])

  const handlePurchase = async () => {
    // TODO: llamar purchaseTicket({ event_id: id })
    // TODO: navigate('/purchase-success') en caso de éxito
    // TODO: mostrar error si no hay cupo o falla la request
  }

  if (loading) return <LoadingSpinner fullScreen />

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: renderizar imagen, título, descripción, fecha, hora, cupo disponible */}
        {/* TODO: mostrar botón "Comprar entrada" si el usuario está autenticado */}
      </main>
    </div>
  )
}

export default EventDetailPage
