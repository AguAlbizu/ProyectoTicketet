import { useState, useEffect, useCallback } from 'react'
import Navbar from '../../components/common/Navbar'
import Footer from '../../components/common/Footer'
import LoadingSpinner from '../../components/common/LoadingSpinner'
import EventFormModal from './EventFormModal'
import EventReportModal from './EventReportModal'
import CreateAdminModal from './CreateAdminModal'
import * as adminApi from '../../api/adminApi'

function AdminEventsPage() {
  const [events, setEvents] = useState([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  const [showForm, setShowForm] = useState(false)
  const [editingEvent, setEditingEvent] = useState(null)
  const [reportEventId, setReportEventId] = useState(null)
  const [showCreateAdmin, setShowCreateAdmin] = useState(false)

  const fetchEvents = useCallback(async () => {
    try {
      setLoading(true)
      setError('')
      const res = await adminApi.getAllEventsAdmin()
      setEvents(res.data)
    } catch {
      setError('Error al cargar los eventos')
    } finally {
      setLoading(false)
    }
  }, [])

  useEffect(() => { fetchEvents() }, [fetchEvents])

  const handleCreateClick = () => {
    setEditingEvent(null)
    setShowForm(true)
  }

  const handleEditClick = (event) => {
    setEditingEvent(event)
    setShowForm(true)
  }

  const handleCancelEvent = async (event) => {
    if (!window.confirm(`¿Cancelar "${event.titulo}"? Se cancelarán todas las entradas activas de este evento.`)) return
    try {
      await adminApi.cancelEvent(event.id_events)
      fetchEvents()
    } catch (err) {
      alert(err.response?.data?.error || 'Error al cancelar el evento')
    }
  }

  const handleFormSave = () => {
    setShowForm(false)
    fetchEvents()
  }

  const formatDate = (dateStr) => {
    return new Date(dateStr).toLocaleDateString('es-AR', {
      day: '2-digit', month: '2-digit', year: 'numeric',
    })
  }

  const formatPrice = (price) =>
    new Intl.NumberFormat('es-AR', {
      style: 'currency', currency: 'ARS', maximumFractionDigits: 0,
    }).format(price)

  return (
    <div className="page-wrapper">
      <Navbar />
      <main className="page-main">
        <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: '1.75rem' }}>
          <div>
            <h1 className="page-title">Panel de Administración</h1>
            <p className="page-subtitle">Gestión de eventos del catálogo</p>
          </div>
          <div style={{ display: 'flex', gap: '0.75rem' }}>
            <button className="btn btn-outline" onClick={() => setShowCreateAdmin(true)}>
              + Crear Admin
            </button>
            <button className="btn btn-primary" onClick={handleCreateClick}>
              + Nuevo Evento
            </button>
          </div>
        </div>

        {loading && <LoadingSpinner />}
        {error && <div className="alert alert-error">{error}</div>}

        {!loading && !error && (
          <div className="admin-table-wrapper">
            <table className="admin-table">
              <thead>
                <tr>
                  <th>Evento</th>
                  <th>Categoría</th>
                  <th>Fecha</th>
                  <th>Precio</th>
                  <th>Cupo</th>
                  <th>Estado</th>
                  <th>Acciones</th>
                </tr>
              </thead>
              <tbody>
                {events.length === 0 && (
                  <tr>
                    <td colSpan={7} style={{ textAlign: 'center', color: 'var(--text-muted)', padding: '2rem' }}>
                      No hay eventos registrados.
                    </td>
                  </tr>
                )}
                {events.map((event) => (
                  <tr key={event.id_events}>
                    <td style={{ maxWidth: 220 }}>
                      <div style={{ fontWeight: 600, lineHeight: 1.3 }}>{event.titulo}</div>
                    </td>
                    <td>
                      {event.categoria && (
                        <span className="badge badge-category">{event.categoria}</span>
                      )}
                    </td>
                    <td style={{ whiteSpace: 'nowrap' }}>
                      {formatDate(event.fecha)} {event.hora}
                    </td>
                    <td style={{ whiteSpace: 'nowrap' }}>{formatPrice(event.precio)}</td>
                    <td>
                      <span style={{ color: event.cupo_disponible === 0 ? 'var(--error)' : 'inherit', fontWeight: 500 }}>
                        {event.cupo_disponible}/{event.capacidad}
                      </span>
                    </td>
                    <td>
                      <span className={`ticket-status ${event.estado === 'activo' ? 'status-activo' : 'status-cancelado'}`}>
                        {event.estado}
                      </span>
                    </td>
                    <td>
                      <div style={{ display: 'flex', gap: '0.4rem', flexWrap: 'wrap' }}>
                        <button
                          className="btn btn-outline"
                          style={{ fontSize: '0.78rem', padding: '0.3rem 0.65rem' }}
                          onClick={() => setReportEventId(event.id_events)}
                        >
                          Reporte
                        </button>
                        <button
                          className="btn btn-outline"
                          style={{ fontSize: '0.78rem', padding: '0.3rem 0.65rem' }}
                          onClick={() => handleEditClick(event)}
                        >
                          Editar
                        </button>
                        {event.estado === 'activo' && (
                          <button
                            className="btn btn-danger"
                            style={{ fontSize: '0.78rem', padding: '0.3rem 0.65rem' }}
                            onClick={() => handleCancelEvent(event)}
                          >
                            Cancelar
                          </button>
                        )}
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </main>
      <Footer />

      {showForm && (
        <EventFormModal
          event={editingEvent}
          onSave={handleFormSave}
          onClose={() => setShowForm(false)}
        />
      )}

      {reportEventId && (
        <EventReportModal
          eventId={reportEventId}
          onClose={() => setReportEventId(null)}
        />
      )}

      {showCreateAdmin && (
        <CreateAdminModal onClose={() => setShowCreateAdmin(false)} />
      )}
    </div>
  )
}

export default AdminEventsPage
