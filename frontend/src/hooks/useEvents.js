import { useState, useEffect } from 'react'
import { getEvents } from '../api/eventsApi'

// Fetches the list of active events, optionally filtered by category.
// Returns { events, loading, error }.
export function useEvents(category = '') {
  const [events, setEvents] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    // TODO: call getEvents(category)
    // TODO: setEvents(response.data), handle errors with setError
    // TODO: setLoading(false) in finally block
  }, [category])

  return { events, loading, error }
}
