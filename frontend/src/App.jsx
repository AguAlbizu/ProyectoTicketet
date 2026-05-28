import { Routes, Route } from 'react-router-dom'

// Auth pages
import LoginPage from './pages/auth/LoginPage'
import RegisterPage from './pages/auth/RegisterPage'

// Client pages
import HomePage from './pages/client/HomePage'
import EventDetailPage from './pages/client/EventDetailPage'
import MyTicketsPage from './pages/client/MyTicketsPage'
import PurchaseSuccessPage from './pages/client/PurchaseSuccessPage'

// Admin pages
import AdminDashboardPage from './pages/admin/AdminDashboardPage'
import EventFormPage from './pages/admin/EventFormPage'
import EventReportPage from './pages/admin/EventReportPage'
import RaffleManagerPage from './pages/admin/RaffleManagerPage'

// Common components
import ProtectedRoute from './components/common/ProtectedRoute'

function App() {
  return (
    <Routes>
      {/* Public routes */}
      <Route path="/login" element={<LoginPage />} />
      <Route path="/register" element={<RegisterPage />} />
      <Route path="/" element={<HomePage />} />
      <Route path="/events/:id" element={<EventDetailPage />} />

      {/* Client-only routes (role: cliente) */}
      <Route
        path="/my-tickets"
        element={
          <ProtectedRoute requiredRole="cliente">
            <MyTicketsPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/purchase-success"
        element={
          <ProtectedRoute requiredRole="cliente">
            <PurchaseSuccessPage />
          </ProtectedRoute>
        }
      />

      {/* Admin-only routes (role: admin) */}
      <Route
        path="/admin"
        element={
          <ProtectedRoute requiredRole="admin">
            <AdminDashboardPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/admin/events/new"
        element={
          <ProtectedRoute requiredRole="admin">
            <EventFormPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/admin/events/:id/edit"
        element={
          <ProtectedRoute requiredRole="admin">
            <EventFormPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/admin/events/:id/report"
        element={
          <ProtectedRoute requiredRole="admin">
            <EventReportPage />
          </ProtectedRoute>
        }
      />
      <Route
        path="/admin/raffles/:id"
        element={
          <ProtectedRoute requiredRole="admin">
            <RaffleManagerPage />
          </ProtectedRoute>
        }
      />
    </Routes>
  )
}

export default App
