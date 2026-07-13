import { useState, useEffect } from 'react'
import * as adminApi from '../../api/adminApi'
import LoadingSpinner from '../../components/common/LoadingSpinner'

function EventReportModal({ eventId, onClose }) {
  const [report, setReport] = useState(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')

  useEffect(() => {
    adminApi.getEventReport(eventId)
      .then((res) => setReport(res.data))
      .catch(() => setError('Error al cargar el reporte'))
      .finally(() => setLoading(false))
  }, [eventId])

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div
        className="modal"
        onClick={(e) => e.stopPropagation()}
        style={{ maxWidth: 640, maxHeight: '90vh', overflowY: 'auto' }}
      >
        <h2 className="modal-title">Reporte de Evento</h2>

        {loading && <LoadingSpinner />}
        {error && <div className="alert alert-error">{error}</div>}

        {report && (
          <>
            <p className="modal-desc">{report.titulo}</p>

            <div style={{ display: 'grid', gridTemplateColumns: 'repeat(3, 1fr)', gap: '1rem', marginBottom: '1.5rem' }}>
              {[
                { value: report.capacidad, label: 'Capacidad total' },
                { value: report.entradas_vendidas, label: 'Entradas vendidas' },
                { value: `${report.porcentaje_ocupacion.toFixed(1)}%`, label: 'Ocupación' },
              ].map(({ value, label }) => (
                <div
                  key={label}
                  style={{
                    background: 'var(--bg)',
                    borderRadius: 'var(--r-sm)',
                    padding: '1rem',
                    textAlign: 'center',
                    border: '1px solid var(--border)',
                  }}
                >
                  <div style={{ fontSize: '1.5rem', fontWeight: 700, color: 'var(--brand)' }}>{value}</div>
                  <div style={{ fontSize: '0.75rem', color: 'var(--text-muted)', marginTop: '0.25rem' }}>{label}</div>
                </div>
              ))}
            </div>

            <h3 style={{ fontSize: '0.875rem', fontWeight: 600, marginBottom: '0.75rem' }}>
              Compradores activos ({report.compradores.length})
            </h3>

            {report.compradores.length === 0 ? (
              <p style={{ color: 'var(--text-muted)', fontSize: '0.875rem' }}>
                No hay entradas activas para este evento.
              </p>
            ) : (
              <div className="admin-table-wrapper">
                <table className="admin-table">
                  <thead>
                    <tr>
                      <th>Nombre</th>
                      <th>Email</th>
                      <th>Fecha compra</th>
                    </tr>
                  </thead>
                  <tbody>
                    {report.compradores.map((buyer) => (
                      <tr key={buyer.ticket_id}>
                        <td>{buyer.nombre}</td>
                        <td style={{ color: 'var(--text-muted)' }}>{buyer.email}</td>
                        <td style={{ color: 'var(--text-muted)', fontSize: '0.8rem' }}>
                          {new Date(buyer.fecha_compra).toLocaleDateString('es-AR')}
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            )}
          </>
        )}

        <div className="modal-actions" style={{ marginTop: '1.5rem' }}>
          <button className="btn btn-outline" onClick={onClose}>Cerrar</button>
        </div>
      </div>
    </div>
  )
}

export default EventReportModal
