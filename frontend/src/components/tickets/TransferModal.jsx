import { useState } from 'react'
import { transferTicket } from '../../api/ticketsApi'

function TransferModal({ ticketId, isOpen, onClose, onSuccess }) {
  const [targetEmail, setTargetEmail] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  if (!isOpen) return null

  const handleConfirm = async () => {
    if (!targetEmail) {
      setError('Ingresá el email del destinatario')
      return
    }
    setError('')
    setLoading(true)
    try {
      await transferTicket(ticketId, targetEmail)
      setTargetEmail('')
      onSuccess()
    } catch (err) {
      setError(err.response?.data?.error || 'Error al transferir la entrada')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={(e) => e.stopPropagation()}>
        <h3 className="modal-title">Transferir entrada</h3>
        <p className="modal-desc">
          Ingresá el email del usuario que recibirá esta entrada. La operación es irreversible.
        </p>
        <div className="form-group" style={{ marginBottom: 0 }}>
          <label className="form-label">Email del destinatario</label>
          <input
            className="form-input"
            type="email"
            value={targetEmail}
            onChange={(e) => setTargetEmail(e.target.value)}
            placeholder="email@ejemplo.com"
            autoFocus
          />
        </div>
        {error && <div className="alert alert-error" style={{ marginTop: '0.75rem', marginBottom: 0 }}>{error}</div>}
        <div className="modal-actions">
          <button className="btn btn-ghost" onClick={onClose} disabled={loading}>Cancelar</button>
          <button className="btn btn-primary" onClick={handleConfirm} disabled={loading}>
            {loading ? 'Transfiriendo...' : 'Confirmar transferencia'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default TransferModal
