import { useState, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { getEvents, cancelEvent } from '../../api/eventsApi'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'

// Admin dashboard — lists all events with options to edit, cancel, view report, and manage raffle.
function AdminDashboardPage() {
  const [events, setEvents] = useState([])
  const [loading, setLoading] = useState(true)
  const navigate = useNavigate()

  useEffect(() => {
    // TODO: call getEvents() (no filter, admin sees all)
    // TODO: setEvents(response.data), setLoading(false)
  }, [])

  const handleCancel = async (eventId) => {
    // TODO: call cancelEvent(eventId) with confirmation dialog first
    // TODO: update local state: remove or mark event as cancelled
  }

  if (loading) return <LoadingSpinner fullScreen />

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: render "Crear evento" button → navigate to /admin/events/new */}
        {/* TODO: render events table with columns: title, date, capacity, available, status */}
        {/* TODO: action buttons per row: Edit, Report, Raffle, Cancel */}
      </main>
    </div>
  )
}

export default AdminDashboardPage
