const CATEGORIAS = ['Todas', 'Música', 'Teatro', 'Deportes', 'Cine', 'Otro']

function EventFilter({ onFilter }) {
  const handleChange = (e) => {
    const value = e.target.value === 'Todas' ? '' : e.target.value
    onFilter(value)
  }

  return (
    <div style={{ marginBottom: '1rem', display: 'flex', gap: '1rem', alignItems: 'center' }}>
      <label htmlFor="categoria">Filtrar por categoría:</label>
      <select id="categoria" onChange={handleChange}>
        {CATEGORIAS.map((cat) => (
          <option key={cat} value={cat}>{cat}</option>
        ))}
      </select>
    </div>
  )
}

export default EventFilter
