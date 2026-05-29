const estadoColor = {
  activo: 'green',
  cancelado: 'red',
  transferido: 'gray',
}

function TicketCard({ ticket, onCancel, onTransfer }) {
  const fecha = ticket.event?.fecha
    ? new Date(ticket.event.fecha).toLocaleDateString('es-AR')
    : ''

  return (
    <div style={{ border: '1px solid #ccc', borderRadius: '8px', padding: '1rem', marginBottom: '1rem' }}>
      <h3>{ticket.event?.titulo || `Evento #${ticket.event_id}`}</h3>
      <p>{fecha} — {ticket.event?.hora}</p>
      <p>
        Estado:{' '}
        <strong style={{ color: estadoColor[ticket.estado] || 'black' }}>
          {ticket.estado}
        </strong>
      </p>

      {ticket.estado === 'activo' && (
        <div style={{ display: 'flex', gap: '0.5rem', marginTop: '0.5rem' }}>
          <button onClick={() => onCancel(ticket.id)}>Cancelar</button>
          <button onClick={() => onTransfer(ticket.id)}>Transferir</button>
        </div>
      )}
    </div>
  )
}

export default TicketCard
