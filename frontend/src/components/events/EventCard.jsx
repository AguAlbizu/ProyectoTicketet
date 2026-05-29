function EventCard({ event, onClick }) {
  const fecha = event.fecha
    ? new Date(event.fecha).toLocaleDateString('es-AR')
    : ''

  return (
    <div
      onClick={() => onClick(event.id)}
      style={{ border: '1px solid #ccc', borderRadius: '8px', padding: '1rem', cursor: 'pointer' }}
    >
      {event.imagen_url && (
        <img
          src={event.imagen_url}
          alt={event.titulo}
          style={{ width: '100%', height: '180px', objectFit: 'cover', borderRadius: '4px' }}
        />
      )}
      <h3>{event.titulo}</h3>
      <p>{fecha} — {event.hora}</p>
      <p><strong>Categoría:</strong> {event.categoria}</p>
      {event.cupo_disponible === 0 ? (
        <span style={{ color: 'red', fontWeight: 'bold' }}>Sin cupo</span>
      ) : (
        <span style={{ color: 'green' }}>Cupo disponible: {event.cupo_disponible}</span>
      )}
    </div>
  )
}

export default EventCard
