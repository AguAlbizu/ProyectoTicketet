import { useNavigate } from 'react-router-dom'
import Navbar from '../../components/common/Navbar'

function PurchaseSuccessPage() {
  const navigate = useNavigate()

  return (
    <div>
      <Navbar />
      <main style={{ textAlign: 'center', padding: '4rem 2rem' }}>
        <h1>¡Entrada comprada con éxito!</h1>
        <p>Tu entrada fue reservada correctamente.</p>
        <div style={{ display: 'flex', gap: '1rem', justifyContent: 'center', marginTop: '2rem' }}>
          <button onClick={() => navigate('/my-tickets')} style={{ padding: '0.75rem 1.5rem' }}>
            Ver mis entradas
          </button>
          <button onClick={() => navigate('/')} style={{ padding: '0.75rem 1.5rem' }}>
            Volver al catálogo
          </button>
        </div>
      </main>
    </div>
  )
}

export default PurchaseSuccessPage
