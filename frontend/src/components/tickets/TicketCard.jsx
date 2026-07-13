import SorteoPanel from '../sorteos/SorteoPanel'

const STRIPE_CLASS = {
  activo:      'ticket-stripe',
  cancelado:   'ticket-stripe ticket-stripe-cancelado',
  transferido: 'ticket-stripe ticket-stripe-transferido',
}

const STATUS_CLASS = {
  activo:      'ticket-status status-activo',
  cancelado:   'ticket-status status-cancelado',
  transferido: 'ticket-status status-transferido',
}

function TicketCard({ ticket, onCancel, onTransfer }) {
  const fecha = ticket.event?.fecha
    ? new Date(ticket.event.fecha).toLocaleDateString('es-AR', { day: 'numeric', month: 'short', year: 'numeric' })
    : ''

  const stripeClass = STRIPE_CLASS[ticket.estado] || 'ticket-stripe'
  const statusClass = STATUS_CLASS[ticket.estado] || 'ticket-status'

  return (
    <div className="ticket-card">
      <div className={stripeClass} />
      <div className="ticket-content">
        <h3 className="ticket-title">
          {ticket.event?.titulo || `Evento #${ticket.id_events}`}
        </h3>
        {fecha && (
          <p className="ticket-meta">{fecha}{ticket.event?.hora ? ` — ${ticket.event.hora}` : ''}</p>
        )}
        <div className="ticket-footer">
          <span className={statusClass}>{ticket.estado}</span>
          {ticket.estado === 'activo' && (
            <div className="ticket-actions">
              <button className="btn btn-danger" onClick={() => onCancel(ticket.id_tickets)}>
                Cancelar
              </button>
              <button className="btn btn-outline" onClick={() => onTransfer(ticket.id_tickets)}>
                Transferir
              </button>
            </div>
          )}
        </div>
        {ticket.estado === 'activo' && (
          <SorteoPanel eventId={ticket.id_events} compact />
        )}
      </div>
    </div>
  )
}

export default TicketCard
