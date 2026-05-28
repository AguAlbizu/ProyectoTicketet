import { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { getEventById } from '../../api/eventsApi'
import { getRaffleByEvent } from '../../api/raffleApi'
import { purchaseTicket } from '../../api/ticketsApi'
import RaffleCard from '../../components/raffle/RaffleCard'
import BuyChancesModal from '../../components/raffle/BuyChancesModal'
import { buyChances } from '../../api/raffleApi'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'
import { useAuth } from '../../hooks/useAuth'

// Event detail page — shows full event info, purchase button, and raffle section.
function EventDetailPage() {
  const { id } = useParams()
  const { user } = useAuth()
  const navigate = useNavigate()

  const [event, setEvent] = useState(null)
  const [raffle, setRaffle] = useState(null)
  const [loading, setLoading] = useState(true)
  const [raffleModalOpen, setRaffleModalOpen] = useState(false)

  useEffect(() => {
    // TODO: fetch event with getEventById(id)
    // TODO: fetch raffle with getRaffleByEvent(id), handle 404 gracefully
    // TODO: setLoading(false) in finally
  }, [id])

  const handlePurchase = async () => {
    // TODO: call purchaseTicket({ event_id: id })
    // TODO: navigate to /purchase-success on success
    // TODO: show error toast on failure
  }

  const handleBuyChances = async (raffleId, quantity) => {
    // TODO: call buyChances(raffleId, { quantity })
    // TODO: refresh raffle/entry state on success
    // TODO: close modal: setRaffleModalOpen(false)
  }

  if (loading) return <LoadingSpinner fullScreen />

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: render event image, title, description, date, capacity */}
        {/* TODO: show "Comprar entrada" button (visible to authenticated clients) */}
        {/* TODO: show RaffleCard if raffle exists */}
        {raffle && (
          <RaffleCard
            raffle={raffle}
            userEntry={null}
            onBuyChances={() => setRaffleModalOpen(true)}
          />
        )}
        <BuyChancesModal
          raffle={raffle}
          isOpen={raffleModalOpen}
          onClose={() => setRaffleModalOpen(false)}
          onConfirm={handleBuyChances}
        />
      </main>
    </div>
  )
}

export default EventDetailPage
