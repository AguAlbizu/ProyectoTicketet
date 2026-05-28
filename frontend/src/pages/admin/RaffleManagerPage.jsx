import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import { getRaffleByEvent, drawRaffle, createRaffle } from '../../api/raffleApi'
import Navbar from '../../components/common/Navbar'
import LoadingSpinner from '../../components/common/LoadingSpinner'

// Admin page to manage a raffle: view entries, create raffle, and execute the draw.
function RaffleManagerPage() {
  const { id } = useParams() // raffle ID
  const [raffle, setRaffle] = useState(null)
  const [loading, setLoading] = useState(true)
  const [drawing, setDrawing] = useState(false)

  useEffect(() => {
    // TODO: fetch raffle by ID (consider GET /api/raffles/:id endpoint)
    // TODO: setRaffle(response.data), setLoading(false)
  }, [id])

  const handleDraw = async () => {
    // TODO: show confirmation dialog before executing
    // TODO: setDrawing(true), call drawRaffle(id)
    // TODO: update raffle state with winner, setDrawing(false)
  }

  if (loading) return <LoadingSpinner fullScreen />

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: render raffle name, event name, status */}
        {/* TODO: render list of participants with their chance counts */}
        {/* TODO: if status === "pendiente": show "Ejecutar sorteo" button */}
        {/* TODO: if status === "realizado": show winner name and email */}
      </main>
    </div>
  )
}

export default RaffleManagerPage
