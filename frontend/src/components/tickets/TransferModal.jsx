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
    <div style={{ position: 'fixed', inset: 0, background: 'rgba(0,0,0,0.5)', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: 'white', padding: '2rem', borderRadius: '8px', minWidth: '320px' }}>
        <h3>Transferir entrada</h3>
        <p>Ingresá el email del usuario que recibirá la entrada:</p>
        <input
          type="email"
          value={targetEmail}
          onChange={(e) => setTargetEmail(e.target.value)}
          placeholder="email@ejemplo.com"
          style={{ width: '100%', padding: '0.5rem', marginBottom: '0.5rem' }}
        />
        {error && <p style={{ color: 'red' }}>{error}</p>}
        <div style={{ display: 'flex', gap: '0.5rem', justifyContent: 'flex-end' }}>
          <button onClick={onClose} disabled={loading}>Cancelar</button>
          <button onClick={handleConfirm} disabled={loading}>
            {loading ? 'Transfiriendo...' : 'Confirmar'}
          </button>
        </div>
      </div>
    </div>
  )
}

export default TransferModal
