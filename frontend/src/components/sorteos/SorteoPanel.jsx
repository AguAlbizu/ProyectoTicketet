import { useState, useEffect, useCallback } from 'react'
import { getSorteoByEvent, buyChances, getMyChances } from '../../api/sorteosApi'
import { useAuth } from '../../hooks/useAuth'

// SorteoPanel muestra el sorteo asociado a un evento (si tiene uno cargado) y permite
// comprar chances para participar. Se usa tanto en el detalle del evento como en Mis Entradas.
function SorteoPanel({ eventId, compact = false }) {
  const { isAuthenticated } = useAuth()
  const [sorteo, setSorteo] = useState(null)
  const [myChances, setMyChances] = useState(0)
  const [cantidad, setCantidad] = useState(1)
  const [loading, setLoading] = useState(true)
  const [buying, setBuying] = useState(false)
  const [error, setError] = useState('')
  const [success, setSuccess] = useState('')

  const fetchMyChances = useCallback(async (sorteoId) => {
    if (!isAuthenticated) return
    try {
      const res = await getMyChances(sorteoId)
      setMyChances(res.data?.chances || 0)
    } catch {
      // silencioso: si falla, simplemente no se muestra el contador
    }
  }, [isAuthenticated])

  useEffect(() => {
    let active = true
    const fetchSorteo = async () => {
      setLoading(true)
      try {
        const res = await getSorteoByEvent(eventId)
        if (!active) return
        setSorteo(res.data)
        fetchMyChances(res.data.id_sorteo)
      } catch {
        if (active) setSorteo(null)
      } finally {
        if (active) setLoading(false)
      }
    }
    fetchSorteo()
    return () => { active = false }
  }, [eventId, fetchMyChances])

  if (loading || !sorteo) return null

  const handleBuy = async () => {
    setError('')
    setSuccess('')
    setBuying(true)
    try {
      await buyChances(sorteo.id_sorteo, Number(cantidad))
      setSuccess(`¡Sumaste ${cantidad} chance${cantidad > 1 ? 's' : ''}!`)
      setCantidad(1)
      fetchMyChances(sorteo.id_sorteo)
    } catch (err) {
      setError(err.response?.data?.error || 'Error al comprar chances')
    } finally {
      setBuying(false)
    }
  }

  return (
    <div className={`sorteo-panel${compact ? ' sorteo-compact' : ''}`}>
      <div className="sorteo-header">
        <span className="sorteo-title">🎁 {sorteo.nombre}</span>
        <span className="badge badge-sorteo">
          {sorteo.estado === 'activo' ? 'Sorteo activo' : 'Sorteo finalizado'}
        </span>
      </div>
      <p className="sorteo-desc">
        Cada chance cuesta ${sorteo.valor_chance?.toLocaleString('es-AR')}. A más chances, más probabilidades de ganar.
      </p>

      {sorteo.estado !== 'activo' && (
        <p className="sorteo-desc">
          {sorteo.ganador ? `Ganador/a: ${sorteo.ganador.nombre}` : 'Este sorteo ya fue realizado.'}
        </p>
      )}

      {sorteo.estado === 'activo' && !isAuthenticated && (
        <p className="sorteo-desc">Iniciá sesión y comprá una entrada para participar.</p>
      )}

      {sorteo.estado === 'activo' && isAuthenticated && (
        <>
          {error && <div className="alert alert-error">{error}</div>}
          {success && <div className="alert alert-success">{success}</div>}
          <div className="sorteo-buy-row">
            <input
              type="number"
              min="1"
              className="form-input sorteo-qty-input"
              value={cantidad}
              onChange={(e) => setCantidad(Math.max(1, Number(e.target.value)))}
            />
            <button className="btn btn-primary" onClick={handleBuy} disabled={buying}>
              {buying ? 'Comprando...' : 'Comprar chances'}
            </button>
          </div>
          {myChances > 0 && (
            <p className="sorteo-my-chances">Ya tenés {myChances} chance{myChances > 1 ? 's' : ''} cargada{myChances > 1 ? 's' : ''}.</p>
          )}
        </>
      )}
    </div>
  )
}

export default SorteoPanel
