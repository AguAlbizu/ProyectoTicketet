// Displays a single ticket with its event details and available actions.
// Props:
//   ticket   — { id, status, purchase_date, event: { title, date } }
//   onCancel — callback(ticketId) to cancel the ticket
//   onTransfer — callback(ticketId) to open the transfer modal
function TicketCard({ ticket, onCancel, onTransfer }) {
  // TODO: render event name, purchase date, status badge
  // TODO: if status === "activo": show Cancel button (calls onCancel(ticket.id))
  // TODO: if status === "activo": show Transfer button (calls onTransfer(ticket.id))

  return <div>{/* TODO: implement TicketCard UI */}</div>
}

export default TicketCard
