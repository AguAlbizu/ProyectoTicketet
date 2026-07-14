import { useState, useEffect, useRef, useCallback } from 'react'
import { getMyNotifications, markNotificationAsRead, markAllNotificationsAsRead } from '../../api/notificationsApi'

const POLL_INTERVAL_MS = 25000

const ICONS = {
  chance_comprada: '🎟️',
  sorteo_ganador: '🏆',
  sorteo_perdedor: '🎁',
}

function timeAgo(dateStr) {
  const diffMin = Math.floor((Date.now() - new Date(dateStr).getTime()) / 60000)
  if (diffMin < 1) return 'ahora'
  if (diffMin < 60) return `hace ${diffMin} min`
  const diffH = Math.floor(diffMin / 60)
  if (diffH < 24) return `hace ${diffH} h`
  return `hace ${Math.floor(diffH / 24)} d`
}

// NotificationBell muestra las notificaciones in-app del usuario (ej. participación
// y resultado de sorteos) en un panel desplegable desde el ícono de campana del Navbar.
function NotificationBell() {
  const [notifications, setNotifications] = useState([])
  const [open, setOpen] = useState(false)
  const containerRef = useRef(null)

  const fetchNotifications = useCallback(async () => {
    try {
      const res = await getMyNotifications()
      setNotifications(res.data || [])
    } catch {
      // silencioso: si falla, simplemente no se actualiza el listado
    }
  }, [])

  useEffect(() => {
    fetchNotifications()
    const interval = setInterval(fetchNotifications, POLL_INTERVAL_MS)
    return () => clearInterval(interval)
  }, [fetchNotifications])

  useEffect(() => {
    const handleClickOutside = (e) => {
      if (containerRef.current && !containerRef.current.contains(e.target)) {
        setOpen(false)
      }
    }
    document.addEventListener('mousedown', handleClickOutside)
    return () => document.removeEventListener('mousedown', handleClickOutside)
  }, [])

  const unreadCount = notifications.filter((n) => !n.leida).length

  const handleItemClick = async (notification) => {
    if (notification.leida) return
    setNotifications((prev) =>
      prev.map((n) => (n.id_notification === notification.id_notification ? { ...n, leida: true } : n))
    )
    try {
      await markNotificationAsRead(notification.id_notification)
    } catch {
      fetchNotifications()
    }
  }

  const handleMarkAllRead = async () => {
    setNotifications((prev) => prev.map((n) => ({ ...n, leida: true })))
    try {
      await markAllNotificationsAsRead()
    } catch {
      fetchNotifications()
    }
  }

  return (
    <div className="notification-bell" ref={containerRef}>
      <button
        className="notification-bell-btn"
        onClick={() => setOpen((o) => !o)}
        aria-label="Notificaciones"
      >
        🔔
        {unreadCount > 0 && (
          <span className="notification-badge">{unreadCount > 9 ? '9+' : unreadCount}</span>
        )}
      </button>

      {open && (
        <div className="notification-dropdown">
          <div className="notification-dropdown-header">
            <span>Notificaciones</span>
            {unreadCount > 0 && (
              <button className="notification-mark-all" onClick={handleMarkAllRead}>
                Marcar todas como leídas
              </button>
            )}
          </div>

          <div className="notification-list">
            {notifications.length === 0 ? (
              <div className="notification-empty">No tenés notificaciones todavía.</div>
            ) : (
              notifications.map((n) => (
                <div
                  key={n.id_notification}
                  className={`notification-item${n.leida ? '' : ' notification-item-unread'}`}
                  onClick={() => handleItemClick(n)}
                >
                  <span className="notification-item-icon">{ICONS[n.tipo] || '🔔'}</span>
                  <div className="notification-item-body">
                    <div className="notification-item-title">{n.titulo}</div>
                    <div className="notification-item-message">{n.mensaje}</div>
                    <div className="notification-item-time">{timeAgo(n.created_at)}</div>
                  </div>
                  {!n.leida && <span className="notification-item-dot" />}
                </div>
              ))
            )}
          </div>
        </div>
      )}
    </div>
  )
}

export default NotificationBell
