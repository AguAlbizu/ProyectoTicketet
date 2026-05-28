import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { getEventById } from '../../api/eventsApi'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'

// Admin report page — shows ticket sales summary for a specific event.
function EventReportPage() {
  const { id } = useParams()
  const [event, setEvent] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    // TODO: fetch event details and ticket stats
    // TODO: consider a dedicated report endpoint: GET /api/events/:id/report
    // TODO: setEvent, setLoading(false)
  }, [id])

  if (loading) return <LoadingSpinner fullScreen />

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: render event name and date */}
        {/* TODO: render metrics: total capacity, tickets sold, tickets available, revenue */}
        {/* TODO: optional: ticket status breakdown (active / cancelled / transferred) */}
      </main>
    </div>
  )
}

export default EventReportPage
