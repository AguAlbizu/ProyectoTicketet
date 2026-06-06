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
    <div>
      <Navbar />
      <main style={{ padding: '1rem 2rem' }}>
        <h1>Eventos disponibles</h1>
        <EventFilter onFilter={filterByCategoria} />

        {loading && <LoadingSpinner />}
        {error && <p style={{ color: 'red' }}>{error}</p>}
        {!loading && !error && events.length === 0 && (
          <p>No hay eventos disponibles en este momento.</p>
        )}

        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(260px, 1fr))', gap: '1rem' }}>
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
