import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getEventById, createEvent, updateEvent } from '../../api/eventsApi'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'

// Admin form to create or edit an event.
// When :id is present in the URL, it loads the event for editing.
function EventFormPage() {
  const { id } = useParams()
  const navigate = useNavigate()
  const isEditing = Boolean(id)

  const [loading, setLoading] = useState(isEditing)
  const [form, setForm] = useState({
    title: '',
    description: '',
    date: '',
    duration_minutes: '',
    capacity: '',
    category: '',
    image_url: '',
  })

  useEffect(() => {
    if (!isEditing) return
    // TODO: call getEventById(id) to pre-fill form fields
    // TODO: setForm with event data, setLoading(false)
  }, [id, isEditing])

  const handleSubmit = async (e) => {
    e.preventDefault()
    // TODO: if isEditing call updateEvent(id, form), else call createEvent(form)
    // TODO: navigate to /admin on success
    // TODO: show error messages on failure
  }

  if (loading) return <LoadingSpinner fullScreen />

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: render form with controlled inputs for each field in `form` state */}
        {/* TODO: submit button text: isEditing ? "Guardar cambios" : "Crear evento" */}
      </main>
    </div>
  )
}

export default EventFormPage
