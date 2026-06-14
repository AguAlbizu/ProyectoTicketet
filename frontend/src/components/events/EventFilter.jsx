const CATEGORIAS = ['Todas', 'Música', 'Teatro', 'Deportes', 'Cine', 'Otro']

function EventFilter({ onFilter }) {
  const handleChange = (e) => {
    onFilter(e.target.value === 'Todas' ? '' : e.target.value)
  }

  return (
    <div className="filter-bar">
      <span className="filter-label">Filtrar por:</span>
      <select className="form-select" style={{ width: 'auto' }} id="categoria" onChange={handleChange}>
        {CATEGORIAS.map((cat) => (
          <option key={cat} value={cat}>{cat}</option>
        ))}
      </select>
    </div>
  )
}

export default EventFilter
