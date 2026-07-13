import { Routes, Route } from 'react-router-dom'

import LoginPage from './pages/auth/LoginPage'
import RegisterPage from './pages/auth/RegisterPage'
import HomePage from './pages/client/HomePage'
import EventDetailPage from './pages/client/EventDetailPage'
import MyTicketsPage from './pages/client/MyTicketsPage'
import PurchaseSuccessPage from './pages/client/PurchaseSuccessPage'
import AdminEventsPage from './pages/admin/AdminEventsPage'
import ProtectedRoute from './components/common/ProtectedRoute'

function App() {
  return (
    <Routes>
      {/* Rutas públicas */}
      <Route path="/" element={<HomePage />} />
      <Route path="/events/:id" element={<EventDetailPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />

      {/* Rutas protegidas — requieren JWT válido */}
      <Route element={<ProtectedRoute />}>
        <Route path="/my-tickets" element={<MyTicketsPage />} />
        <Route path="/purchase-success" element={<PurchaseSuccessPage />} />
      </Route>

      {/* Rutas de administrador — requieren rol administrador */}
      <Route element={<ProtectedRoute requiredRole="administrador" />}>
        <Route path="/admin/events" element={<AdminEventsPage />} />
      </Route>
    </Routes>
  )
}

export default App
