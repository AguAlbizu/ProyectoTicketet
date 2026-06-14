import { useNavigate } from 'react-router-dom'
import { useEvents } from '../../hooks/useEvents'
import EventCard from '../../components/events/EventCard'
import EventFilter from '../../components/events/EventFilter'
import Navbar from '../../components/common/Navbar'
import Footer from '../../components/common/Footer'
import LoadingSpinner from '../../components/common/LoadingSpinner'

function HomePage() {
  const { events, loading, error, filterByCategoria } = useEvents()
  const navigate = useNavigate()

  return (
    <div className="page-wrapper">
      <Navbar />
      <main className="page-main">
        <h1 className="page-title">Eventos disponibles</h1>
        <p className="page-subtitle">Encontrá tu próxima experiencia y conseguí tus entradas</p>

        <EventFilter onFilter={filterByCategoria} />

        {loading && <LoadingSpinner />}
        {error && <div className="alert alert-error">{error}</div>}

        {!loading && !error && events.length === 0 && (
          <div className="empty-state">
            <div className="empty-state-icon">🎪</div>
            <p className="empty-state-text">No hay eventos disponibles en este momento.</p>
          </div>
        )}

        <div className="events-grid">
          {events.map((event) => (
            <EventCard
              key={event.id_events}
              event={event}
              onClick={(id) => navigate(`/events/${id}`)}
            />
          ))}
        </div>
      </main>
      <Footer />
    </div>
  )
}

export default HomePage
