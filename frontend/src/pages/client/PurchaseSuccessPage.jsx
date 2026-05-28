import { Link } from 'react-router-dom'
import Navbar from '../../components/common/Navbar'

// Confirmation page shown after a successful ticket purchase.
function PurchaseSuccessPage() {
  // TODO: display a success message with the purchased event name (from navigation state or query params)
  // TODO: show links to "Ver mis entradas" (/my-tickets) and "Volver al catálogo" (/)

  return (
    <div>
      <Navbar />
      <main>
        {/* TODO: implement PurchaseSuccessPage UI */}
        <h1>¡Compra exitosa!</h1>
        <Link to="/my-tickets">Ver mis entradas</Link>
        <Link to="/">Volver al catálogo</Link>
      </main>
    </div>
  )
}

export default PurchaseSuccessPage
