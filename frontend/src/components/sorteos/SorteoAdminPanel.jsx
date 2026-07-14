import { useState, useEffect, useCallback } from 'react'
import { getSorteosByEvent, getSorteoChanceSummary, createSorteo, runSorteoDraw } from '../../api/sorteosApi'

// SorteoAdminPanel permite al administrador gestionar el historial de sorteos de un evento:
// crear uno nuevo (incluso después de que otro ya se haya realizado), ejecutar el activo, y ver
// sus participantes con la cantidad de chances que compró cada uno. Se usa dentro del reporte
// de evento del panel admin, en la pestaña "Sorteo".
function SorteoAdminPanel({ eventId }) {
  const [sorteos, setSorteos] = useState([])
  const [participantes, setParticipantes] = useState([])
  const [loading, setLoading] = useState(true)
  const [nombre, setNombre] = useState('')
  const [valorChance, setValorChance] = useState('')
  const [saving, setSaving] = useState(false)
  const [drawing, setDrawing] = useState(false)
  const [error, setError] = useState('')

  const sorteoActivo = sorteos.find((s) => s.estado === 'activo') || null

  const fetchSorteos = useCallback(async () => {
    setLoading(true)
    try {
      const res = await getSorteosByEvent(eventId)
      setSorteos(res.data || [])
    } catch {
      setSorteos([])
    } finally {
      setLoading(false)
    }
  }, [eventId])

  useEffect(() => { fetchSorteos() }, [fetchSorteos])

  useEffect(() => {
    if (!sorteoActivo) {
      setParticipantes([])
      return
    }
    getSorteoChanceSummary(sorteoActivo.id_sorteo)
      .then((res) => setParticipantes(res.data || []))
      .catch(() => setParticipantes([]))
  }, [sorteoActivo])

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
      fetchSorteos()
    } catch (err) {
      setError(err.response?.data?.error || 'Error al crear el sorteo')
    } finally {
      setSaving(false)
    }
  }

  const handleDraw = async () => {
    if (!window.confirm(`¿Realizar el sorteo "${sorteoActivo.nombre}"? Esta acción no se puede deshacer.`)) return
    setError('')
    setDrawing(true)
    try {
      await runSorteoDraw(sorteoActivo.id_sorteo)
      fetchSorteos()
    } catch (err) {
      setError(err.response?.data?.error || 'Error al realizar el sorteo')
    } finally {
      setDrawing(false)
    }
  }

  if (loading) return null

  return (
    <div>
      <div className="sorteo-panel">
        <div className="sorteo-header">
          <span className="sorteo-title">🎁 Sorteo del evento</span>
          {sorteoActivo && <span className="badge badge-sorteo">Activo</span>}
        </div>

        {error && <div className="alert alert-error">{error}</div>}

        {!sorteoActivo && (
          <>
            <p className="sorteo-desc">
              {sorteos.length > 0
                ? 'El último sorteo ya fue realizado. Podés cargar uno nuevo para este evento.'
                : 'Este evento todavía no tiene un sorteo cargado.'}
            </p>
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

        {sorteoActivo && (
          <>
            <p className="sorteo-desc">
              "{sorteoActivo.nombre}" — ${sorteoActivo.valor_chance?.toLocaleString('es-AR')} por chance.
            </p>
            <button className="btn btn-danger" onClick={handleDraw} disabled={drawing || participantes.length === 0}>
              {drawing ? 'Realizando sorteo...' : 'Realizar sorteo'}
            </button>

            <h4 style={{ fontSize: '0.8rem', fontWeight: 600, margin: '1rem 0 0.5rem' }}>
              Participantes ({participantes.length})
            </h4>
            {participantes.length === 0 ? (
              <p style={{ color: 'var(--text-muted)', fontSize: '0.8rem' }}>
                Todavía nadie compró chances para este sorteo.
              </p>
            ) : (
              <div className="admin-table-wrapper">
                <table className="admin-table">
                  <thead>
                    <tr>
                      <th>Nombre</th>
                      <th>Email</th>
                      <th>Chances</th>
                    </tr>
                  </thead>
                  <tbody>
                    {participantes.map((p) => (
                      <tr key={p.user_id}>
                        <td>{p.nombre}</td>
                        <td style={{ color: 'var(--text-muted)' }}>{p.email}</td>
                        <td>{p.cantidad}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}
          </>
        )}
      </div>

      {sorteos.length > 0 && (
        <>
          <h3 style={{ fontSize: '0.875rem', fontWeight: 600, margin: '1.5rem 0 0.75rem' }}>
            Historial de sorteos ({sorteos.length})
          </h3>
          <div className="admin-table-wrapper">
            <table className="admin-table">
              <thead>
                <tr>
                  <th>Nombre</th>
                  <th>Estado</th>
                  <th>Ganador/a</th>
                  <th>Fecha</th>
                </tr>
              </thead>
              <tbody>
                {sorteos.map((s) => (
                  <tr key={s.id_sorteo}>
                    <td>{s.nombre}</td>
                    <td>
                      <span className={`ticket-status status-${s.estado === 'activo' ? 'activo' : 'cancelado'}`}>
                        {s.estado === 'activo' ? 'Activo' : 'Realizado'}
                      </span>
                    </td>
                    <td>{s.ganador?.nombre || '—'}</td>
                    <td style={{ color: 'var(--text-muted)', fontSize: '0.8rem' }}>
                      {s.fecha_realizado ? new Date(s.fecha_realizado).toLocaleDateString('es-AR') : '—'}
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </>
      )}
    </div>
  )
}

export default SorteoAdminPanel
