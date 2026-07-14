import { useState, useEffect, useCallback } from 'react'
import { getSorteoByEvent, createSorteo, runSorteoDraw } from '../../api/sorteosApi'

// SorteoAdminSection permite al administrador cargar el sorteo de un evento (si no tiene uno)
// o ejecutarlo (si ya tiene participantes). Se usa dentro del reporte de evento del panel admin.
function SorteoAdminSection({ eventId }) {
  const [sorteo, setSorteo] = useState(null)
  const [loading, setLoading] = useState(true)
  const [nombre, setNombre] = useState('')
  const [valorChance, setValorChance] = useState('')
  const [saving, setSaving] = useState(false)
  const [drawing, setDrawing] = useState(false)
  const [error, setError] = useState('')

  const fetchSorteo = useCallback(() => {
    setLoading(true)
    getSorteoByEvent(eventId)
      .then((res) => setSorteo(res.data))
      .catch(() => setSorteo(null))
      .finally(() => setLoading(false))
  }, [eventId])

  useEffect(() => { fetchSorteo() }, [fetchSorteo])

  const handleCreate = async () => {
    if (!nombre || !valorChance) {
      setError('Completá nombre y valor de la chance')
      return
    }
    setError('')
    setSaving(true)
    try {
      await createSorteo(eventId, nombre, Number(valorChance))
      setNombre('')
      setValorChance('')
      fetchSorteo()
    } catch (err) {
      setError(err.response?.data?.error || 'Error al crear el sorteo')
    } finally {
      setSaving(false)
    }
  }

  const handleDraw = async () => {
    if (!window.confirm(`¿Realizar el sorteo "${sorteo.nombre}"? Esta acción no se puede deshacer.`)) return
    setError('')
    setDrawing(true)
    try {
      await runSorteoDraw(sorteo.id_sorteo)
      fetchSorteo()
    } catch (err) {
      setError(err.response?.data?.error || 'Error al realizar el sorteo')
    } finally {
      setDrawing(false)
    }
  }

  if (loading) return null

  return (
    <div className="sorteo-panel" style={{ marginTop: '1.5rem' }}>
      <div className="sorteo-header">
        <span className="sorteo-title">🎁 Sorteo del evento</span>
        {sorteo && (
          <span className="badge badge-sorteo">
            {sorteo.estado === 'activo' ? 'Activo' : 'Realizado'}
          </span>
        )}
      </div>

      {error && <div className="alert alert-error">{error}</div>}

      {!sorteo && (
        <>
          <p className="sorteo-desc">Este evento todavía no tiene un sorteo cargado.</p>
          <div className="sorteo-buy-row">
            <input
              className="form-input"
              placeholder="Nombre del sorteo"
              value={nombre}
              onChange={(e) => setNombre(e.target.value)}
              style={{ flex: 1, minWidth: 160 }}
            />
            <input
              type="number"
              min="1"
              className="form-input sorteo-qty-input"
              placeholder="Valor"
              value={valorChance}
              onChange={(e) => setValorChance(e.target.value)}
              style={{ width: 100 }}
            />
            <button className="btn btn-primary" onClick={handleCreate} disabled={saving}>
              {saving ? 'Creando...' : 'Crear sorteo'}
            </button>
          </div>
        </>
      )}

      {sorteo && sorteo.estado === 'activo' && (
        <>
          <p className="sorteo-desc">
            "{sorteo.nombre}" — ${sorteo.valor_chance?.toLocaleString('es-AR')} por chance.
          </p>
          <button className="btn btn-danger" onClick={handleDraw} disabled={drawing}>
            {drawing ? 'Realizando sorteo...' : 'Realizar sorteo'}
          </button>
        </>
      )}

      {sorteo && sorteo.estado !== 'activo' && (
        <p className="sorteo-desc">
          Ganador/a: <strong>{sorteo.ganador?.nombre || 'sin datos'}</strong>
          {sorteo.ganador?.email ? ` (${sorteo.ganador.email})` : ''}
        </p>
      )}
    </div>
  )
}

export default SorteoAdminSection
