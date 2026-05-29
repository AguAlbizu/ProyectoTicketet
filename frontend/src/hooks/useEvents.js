import { useState, useEffect } from 'react'
import { getEvents } from '../api/eventsApi'

export function useEvents() {
  const [events, setEvents] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState(null)

  useEffect(() => {
    fetchEvents('')
  }, [])

  const fetchEvents = async (categoria) => {
    setLoading(true)
    setError(null)
    try {
      const res = await getEvents(categoria)
      setEvents(res.data)
    } catch (err) {
      setError('Error al cargar los eventos')
    } finally {
      setLoading(false)
    }
  }

  const filterByCategoria = (categoria) => {
    fetchEvents(categoria)
  }

  return { events, loading, error, filterByCategoria }
}
