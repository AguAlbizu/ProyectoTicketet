import { useState, useEffect } from 'react'
import { getMyTickets, cancelTicket, transferTicket } from '../../api/ticketsApi'
import TicketCard from '../../components/tickets/TicketCard'
import TransferModal from '../../components/tickets/TransferModal'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'

// Client page — lists the authenticated user's tickets with cancel and transfer options.
function MyTicketsPage() {
  const [tickets, setTickets] = useState([])
  const [loading, setLoading] = useState(true)
  const [transferTicketId, setTransferTicketId] = useState(null)

  useEffect(() => {
    // TODO: call getMyTickets()
    // TODO: setTickets(response.data), setLoading(false)
  }, [])

  const handleCancel = async (ticketId) => {
    // TODO: call cancelTicket(ticketId)
    // TODO: update local state: set ticket status to "cancelado"
  }

  const handleTransfer = async (ticketId, targetEmail) => {
    // TODO: call transferTicket(ticketId, { target_email: targetEmail })
    // TODO: update local state, close modal
  }

  if (loading) return <LoadingSpinner fullScreen />

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: render page title */}
        {tickets.map((ticket) => (
          <TicketCard
            key={ticket.id}
            ticket={ticket}
            onCancel={handleCancel}
            onTransfer={(id) => setTransferTicketId(id)}
          />
        ))}
        <TransferModal
          ticketId={transferTicketId}
          isOpen={!!transferTicketId}
          onClose={() => setTransferTicketId(null)}
          onConfirm={handleTransfer}
        />
      </main>
    </div>
  )
}

export default MyTicketsPage
