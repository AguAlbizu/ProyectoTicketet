const CATEGORY_EMOJI = {
  'Música':    '🎵',
  'Teatro':    '🎭',
  'Deportes':  '⚽',
  'Cine':      '🎬',
}

function EventCard({ event, onClick }) {
  const fecha = event.fecha
    ? new Date(event.fecha).toLocaleDateString('es-AR', { day: 'numeric', month: 'short', year: 'numeric' })
    : ''

  const emoji = CATEGORY_EMOJI[event.categoria] || '🎪'
  const sinCupo = event.cupo_disponible === 0

  return (
    <div className="card card-clickable" onClick={() => onClick(event.id_events)}>
      {event.imagen_url ? (
        <img className="card-img" src={event.imagen_url} alt={event.titulo} />
      ) : (
        <div className="card-img-placeholder">{emoji}</div>
      )}
      <div className="card-body">
        <p className="event-card-meta">{fecha} — {event.hora}</p>
        <h3 className="event-card-title">{event.titulo}</h3>
        <div className="event-card-footer">
          <span className="badge badge-category">{event.categoria}</span>
          {sinCupo ? (
            <span className="badge badge-soldout">Sin cupo</span>
          ) : (
            <span className="badge badge-available">{event.cupo_disponible} disp.</span>
          )}
        </div>
        {event.precio > 0 && (
          <p className="event-card-price">
            ${event.precio.toLocaleString('es-AR')}
          </p>
        )}
      </div>
    </div>
  )
}

export default EventCard
