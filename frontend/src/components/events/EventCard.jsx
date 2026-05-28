import { Link } from 'react-router-dom'

// Displays a summary card for a single event.
// Props:
//   event — { id, title, date, category, image_url, available_tickets, status }
function EventCard({ event }) {
  // TODO: render event image, title, date (formatted), category badge
  // TODO: show "Agotado" if available_tickets === 0
  // TODO: wrap card in a <Link to={`/events/${event.id}`}>

  return (
    <div>
      {/* TODO: implement EventCard UI */}
      <Link to={`/events/${event?.id}`}>{event?.title}</Link>
    </div>
  )
}

export default EventCard
