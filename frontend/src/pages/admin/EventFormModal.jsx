import { useState } from 'react'
import * as adminApi from '../../api/adminApi'

const CATEGORIAS = ['Música', 'Teatro', 'Deportes', 'Cine', 'Otro']

function toDateInput(isoStr) {
  if (!isoStr) return ''
  return isoStr.substring(0, 10)
}

function toISODate(dateStr) {
  if (!dateStr) return ''
  return new Date(dateStr + 'T00:00:00Z').toISOString()
}

function EventFormModal({ event, onSave, onClose }) {
  const isEdit = !!event

  const [form, setForm] = useState({
    titulo: event?.titulo || '',
    descripcion: event?.descripcion || '',
    fecha: toDateInput(event?.fecha),
    hora: event?.hora || '',
    capacidad: event?.capacidad || '',
    categoria: event?.categoria || '',
    direccion: event?.direccion || '',
    imagen_url: event?.imagen_url || '',
    precio: event?.precio ?? '',
    estado: event?.estado || 'activo',
  })
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value })
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const payload = {
        titulo: form.titulo,
        descripcion: form.descripcion,
        fecha: toISODate(form.fecha),
        hora: form.hora,
        capacidad: parseInt(form.capacidad),
        categoria: form.categoria,
        direccion: form.direccion,
        imagen_url: form.imagen_url,
        precio: parseInt(form.precio) || 0,
        ...(isEdit && { estado: form.estado }),
      }
      if (isEdit) {
        await adminApi.updateEvent(event.id_events, payload)
      } else {
        await adminApi.createEvent(payload)
      }
      onSave()
    } catch (err) {
      setError(err.response?.data?.error || 'Error al guardar el evento')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div
        className="modal"
        onClick={(e) => e.stopPropagation()}
        style={{ maxWidth: 560, overflowY: 'auto', maxHeight: '90vh' }}
      >
        <h2 className="modal-title">{isEdit ? 'Editar Evento' : 'Nuevo Evento'}</h2>
        <p className="modal-desc">
          {isEdit ? `Editando: ${event.titulo}` : 'Completá los datos del nuevo evento'}
        </p>

        {error && <div className="alert alert-error">{error}</div>}

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label className="form-label">Título *</label>
            <input
              className="form-input"
              name="titulo"
              value={form.titulo}
              onChange={handleChange}
              required
            />
          </div>

          <div className="form-group">
            <label className="form-label">Descripción</label>
            <textarea
              className="form-input"
              name="descripcion"
              value={form.descripcion}
              onChange={handleChange}
              rows={3}
              style={{ resize: 'vertical' }}
            />
          </div>

          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
            <div className="form-group">
              <label className="form-label">Fecha *</label>
              <input
                className="form-input"
                type="date"
                name="fecha"
                value={form.fecha}
                onChange={handleChange}
                required
              />
            </div>
            <div className="form-group">
              <label className="form-label">Hora *</label>
              <input
                className="form-input"
                type="time"
                name="hora"
                value={form.hora}
                onChange={handleChange}
                required
              />
            </div>
          </div>

          <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
            <div className="form-group">
              <label className="form-label">Capacidad *</label>
              <input
                className="form-input"
                type="number"
                name="capacidad"
                value={form.capacidad}
                onChange={handleChange}
                min={1}
                required
              />
            </div>
            <div className="form-group">
              <label className="form-label">Precio (ARS)</label>
              <input
                className="form-input"
                type="number"
                name="precio"
                value={form.precio}
                onChange={handleChange}
                min={0}
              />
            </div>
          </div>

          <div className="form-group">
            <label className="form-label">Categoría</label>
            <select
              className="form-select"
              name="categoria"
              value={form.categoria}
              onChange={handleChange}
            >
              <option value="">Sin categoría</option>
              {CATEGORIAS.map((c) => (
                <option key={c} value={c}>{c}</option>
              ))}
            </select>
          </div>

          <div className="form-group">
            <label className="form-label">Dirección</label>
            <input
              className="form-input"
              name="direccion"
              value={form.direccion}
              onChange={handleChange}
            />
          </div>

          <div className="form-group">
            <label className="form-label">URL de imagen</label>
            <input
              className="form-input"
              name="imagen_url"
              value={form.imagen_url}
              onChange={handleChange}
              placeholder="https://..."
            />
          </div>

          {isEdit && (
            <div className="form-group">
              <label className="form-label">Estado</label>
              <select
                className="form-select"
                name="estado"
                value={form.estado}
                onChange={handleChange}
              >
                <option value="activo">Activo</option>
                <option value="cancelado">Cancelado</option>
              </select>
            </div>
          )}

          <div className="modal-actions">
            <button type="button" className="btn btn-outline" onClick={onClose}>
              Cancelar
            </button>
            <button type="submit" className="btn btn-primary" disabled={loading}>
              {loading ? 'Guardando...' : isEdit ? 'Guardar cambios' : 'Crear evento'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

export default EventFormModal
