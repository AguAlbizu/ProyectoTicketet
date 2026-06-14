import { useNavigate } from 'react-router-dom'
import Navbar from '../../components/common/Navbar'
import Footer from '../../components/common/Footer'

function PurchaseSuccessPage() {
  const navigate = useNavigate()

  return (
    <div className="page-wrapper">
      <Navbar />
      <main className="success-main">
        <div className="success-icon">✓</div>
        <h1 className="success-title">¡Entrada comprada!</h1>
        <p className="success-subtitle">Tu entrada fue reservada correctamente. ¡Disfrutá el evento!</p>
        <div className="success-actions">
          <button className="btn btn-primary btn-lg" onClick={() => navigate('/my-tickets')}>
            Ver mis entradas
          </button>
          <button className="btn btn-outline btn-lg" onClick={() => navigate('/')}>
            Seguir explorando
          </button>
        </div>
      </main>
      <Footer />
    </div>
  )
}

export default PurchaseSuccessPage
