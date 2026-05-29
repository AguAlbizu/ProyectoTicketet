import { Routes, Route } from 'react-router-dom'

// Auth pages
import LoginPage from './pages/auth/LoginPage'
import RegisterPage from './pages/auth/RegisterPage'

// Client pages
import HomePage from './pages/client/HomePage'
import EventDetailPage from './pages/client/EventDetailPage'
import MyTicketsPage from './pages/client/MyTicketsPage'
import PurchaseSuccessPage from './pages/client/PurchaseSuccessPage'

// Common components
import ProtectedRoute from './components/common/ProtectedRoute'

// TODO (entrega final): agregar rutas protegidas por rol de administrador

function App() {
  return (
    <Routes>
      {/* Rutas públicas */}
      <Route path="/" element={<HomePage />} />
      <Route path="/events/:id" element={<EventDetailPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />

      {/* Rutas protegidas — requieren token JWT válido */}
      <Route
        path="/my-tickets"
        element={
          <ProtectedRoute>
            <MyTicketsPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/purchase-success"
        element={
          <ProtectedRoute>
            <PurchaseSuccessPage />
          </ProtectedRoute>
        }
      />
    </Routes>
  )
}

export default App
