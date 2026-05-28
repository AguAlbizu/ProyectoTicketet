import { useState } from 'react'
import { useEvents } from '../../hooks/useEvents'
import EventCard from '../../components/events/EventCard'
import EventFilter from '../../components/events/EventFilter'
import Navbar from '../../components/common/Navbar'
import Footer from '../../components/common/Footer'
import LoadingSpinner from '../../components/common/LoadingSpinner'

// Public home page — shows the full catalog of active events.
function HomePage() {
  const [selectedCategory, setSelectedCategory] = useState('')
  const { events, loading, error } = useEvents(selectedCategory)

  // TODO: derive unique categories list from events for EventFilter
  // TODO: render EventFilter with categories and onChange handler
  // TODO: render EventCard grid from events array
  // TODO: show LoadingSpinner while loading
  // TODO: show error message if error is set

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: implement HomePage layout */}
        {loading && <LoadingSpinner />}
        {error && <p>{error}</p>}
        <EventFilter selectedCategory={selectedCategory} onChange={setSelectedCategory} />
        <div>
          {events.map((event) => (
            <EventCard key={event.id} event={event} />
          ))}
        </div>
      </main>
      <Footer />
    </div>
  )
}

export default HomePage
