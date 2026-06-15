import { useState, useEffect } from 'react'
import { getMyTickets, cancelTicket } from '../../api/ticketsApi'
import TicketCard from '../../components/tickets/TicketCard'
import TransferModal from '../../components/tickets/TransferModal'
import Navbar from '../../components/common/Navbar'
import Footer from '../../components/common/Footer'
import LoadingSpinner from '../../components/common/LoadingSpinner'

const TABS = [
  { key: 'disponibles',  label: 'Disponibles' },
  { key: 'compradas',    label: 'Compradas' },
  { key: 'recibidas',    label: 'Recibidas' },
  { key: 'transferidas', label: 'Transferidas' },
  { key: 'canceladas',   label: 'Canceladas' },
]

function MyTicketsPage() {
  const [tickets, setTickets] = useState([])
  const [loading, setLoading] = useState(true)
  const [activeTab, setActiveTab] = useState('disponibles')
  const [transferTicketId, setTransferTicketId] = useState(null)

  const fetchTickets = async () => {
    setLoading(true)
    try {
      const res = await getMyTickets()
      setTickets(res.data || [])
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => { fetchTickets() }, [])

  const handleCancel = async (ticketId) => {
    if (!window.confirm('¿Estás seguro de que querés cancelar esta entrada?')) return
    try {
      await cancelTicket(ticketId)
      fetchTickets()
    } catch (err) {
      alert(err.response?.data?.error || 'Error al cancelar la entrada')
    }
  }

  // disponibles: todas las entradas activas, sin importar el origen
  const disponibles = tickets.filter(t => t.estado === 'activo')

  // compradas: tickets que este usuario compró originalmente (cualquier estado)
  const compradas = tickets.filter(t => t.origen === 'compra' || !t.origen)

  // recibidas: tickets que llegaron a esta cuenta por transferencia (cualquier estado)
  const recibidas = tickets.filter(t => t.origen === 'transferencia')

  // transferidas: tickets que este usuario transfirió a alguien más
  const transferidas = tickets.filter(t => t.estado === 'transferido')

  // canceladas: tickets cancelados por el usuario
  const canceladas = tickets.filter(t => t.estado === 'cancelado')

  const counts = {
    disponibles:  disponibles.length,
    compradas:    compradas.length,
    recibidas:    recibidas.length,
    transferidas: transferidas.length,
    canceladas:   canceladas.length,
  }

  const currentList = activeTab === 'disponibles'  ? disponibles
    : activeTab === 'compradas'    ? compradas
    : activeTab === 'recibidas'    ? recibidas
    : activeTab === 'canceladas'   ? canceladas
    : transferidas

  const emptyMessages = {
    disponibles:  { icon: '🎟️', text: 'No tenés entradas disponibles en este momento.' },
    compradas:    { icon: '🛒', text: 'Todavía no compraste ninguna entrada.' },
    recibidas:    { icon: '🎁', text: 'No recibiste ninguna entrada por transferencia.' },
    transferidas: { icon: '↗️', text: 'No transferiste ninguna entrada.' },
    canceladas:   { icon: '❌', text: 'No cancelaste ninguna entrada.' },
  }

  return (
    <div className="page-wrapper">
      <Navbar />
      <main className="page-main-narrow">
        <h1 className="page-title">Mis Entradas</h1>
        <p className="page-subtitle">Gestioná tus entradas compradas y recibidas</p>

        {loading && <LoadingSpinner />}

        {!loading && (
          <>
            <div className="tabs">
              {TABS.map(tab => (
                <button
                  key={tab.key}
                  className={`tab-btn${activeTab === tab.key ? ' active' : ''}`}
                  onClick={() => setActiveTab(tab.key)}
                >
                  {tab.label}
                  <span className="tab-count">{counts[tab.key]}</span>
                </button>
              ))}
            </div>

            {currentList.length === 0 ? (
              <div className="empty-state">
                <div className="empty-state-icon">{emptyMessages[activeTab].icon}</div>
                <p className="empty-state-text">{emptyMessages[activeTab].text}</p>
              </div>
            ) : (
              currentList.map(ticket => (
                <TicketCard
                  key={ticket.id_tickets}
                  ticket={ticket}
                  onCancel={handleCancel}
                  onTransfer={(id) => setTransferTicketId(id)}
                />
              ))
            )}
          </>
        )}

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
      <Footer />
    </div>
  )
}

export default MyTicketsPage
