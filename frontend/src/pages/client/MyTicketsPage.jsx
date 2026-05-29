import { useState, useEffect } from 'react'
import { getMyTickets, cancelTicket } from '../../api/ticketsApi'
import TicketCard from '../../components/tickets/TicketCard'
import TransferModal from '../../components/tickets/TransferModal'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'

function MyTicketsPage() {
  const [tickets, setTickets] = useState([])
  const [loading, setLoading] = useState(true)
  const [transferTicketId, setTransferTicketId] = useState(null)

  const fetchTickets = async () => {
    setLoading(true)
    try {
      const res = await getMyTickets()
      setTickets(res.data)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchTickets()
  }, [])

  const handleCancel = async (ticketId) => {
    if (!window.confirm('¿Estás seguro de que querés cancelar esta entrada?')) return
    try {
      await cancelTicket(ticketId)
      fetchTickets()
    } catch (err) {
      alert(err.response?.data?.error || 'Error al cancelar la entrada')
    }
  }

  return (
    <div>
      <Navbar />
      <main style={{ padding: '1rem 2rem', maxWidth: '700px', margin: '0 auto' }}>
        <h1>Mis Entradas</h1>
        {loading && <LoadingSpinner />}
        {!loading && tickets.length === 0 && <p>No tenés entradas aún.</p>}
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
          onSuccess={() => {
            setTransferTicketId(null)
            fetchTickets()
          }}
        />
      </main>
    </div>
  )
}

export default MyTicketsPage
