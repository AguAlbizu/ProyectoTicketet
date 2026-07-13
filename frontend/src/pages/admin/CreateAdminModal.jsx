import { useState } from 'react'
import * as adminApi from '../../api/adminApi'

function CreateAdminModal({ onClose }) {
  const [tab, setTab] = useState('crear')
  const [success, setSuccess] = useState('')

  // Crear nuevo admin
  const [createForm, setCreateForm] = useState({ nombre: '', email: '', password: '' })
  const [createLoading, setCreateLoading] = useState(false)
  const [createError, setCreateError] = useState('')

  // Promover usuario existente
  const [promoteEmail, setPromoteEmail] = useState('')
  const [promoteLoading, setPromoteLoading] = useState(false)
  const [promoteError, setPromoteError] = useState('')

  const handleTabChange = (t) => {
    setTab(t)
    setSuccess('')
    setCreateError('')
    setPromoteError('')
  }

  const handleCreate = async (e) => {
    e.preventDefault()
    setCreateError('')
    setCreateLoading(true)
    try {
      await adminApi.createAdmin(createForm)
      setSuccess('Administrador creado exitosamente.')
    } catch (err) {
      setCreateError(err.response?.data?.error || 'Error al crear el administrador')
    } finally {
      setCreateLoading(false)
    }
  }

  const handlePromote = async (e) => {
    e.preventDefault()
    setPromoteError('')
    setPromoteLoading(true)
    try {
      await adminApi.promoteToAdmin(promoteEmail)
      setSuccess(`El usuario ${promoteEmail} ahora es administrador.`)
    } catch (err) {
      setPromoteError(err.response?.data?.error || 'Error al promover el usuario')
    } finally {
      setPromoteLoading(false)
    }
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal" onClick={(e) => e.stopPropagation()} style={{ maxWidth: 420 }}>
        <h2 className="modal-title">Gestión de Administradores</h2>

        {success ? (
          <>
            <div className="alert alert-success" style={{ marginTop: '1rem' }}>{success}</div>
            <div className="modal-actions">
              <button className="btn btn-primary" onClick={onClose}>Cerrar</button>
            </div>
          </>
        ) : (
          <>
            <div className="tabs" style={{ marginTop: '0.5rem' }}>
              <button
                className={`tab-btn ${tab === 'crear' ? 'active' : ''}`}
                onClick={() => handleTabChange('crear')}
              >
                Crear nuevo
              </button>
              <button
                className={`tab-btn ${tab === 'promover' ? 'active' : ''}`}
                onClick={() => handleTabChange('promover')}
              >
                Promover existente
              </button>
            </div>

            {tab === 'crear' && (
              <form onSubmit={handleCreate}>
                <p className="modal-desc">Crea una cuenta nueva con rol de administrador.</p>
                {createError && <div className="alert alert-error">{createError}</div>}
                <div className="form-group">
                  <label className="form-label">Nombre *</label>
                  <input
                    className="form-input"
                    name="nombre"
                    value={createForm.nombre}
                    onChange={(e) => setCreateForm({ ...createForm, nombre: e.target.value })}
                    required
                  />
                </div>
                <div className="form-group">
                  <label className="form-label">Email *</label>
                  <input
                    className="form-input"
                    type="email"
                    name="email"
                    value={createForm.email}
                    onChange={(e) => setCreateForm({ ...createForm, email: e.target.value })}
                    required
                  />
                </div>
                <div className="form-group">
                  <label className="form-label">Contraseña *</label>
                  <input
                    className="form-input"
                    type="password"
                    name="password"
                    value={createForm.password}
                    onChange={(e) => setCreateForm({ ...createForm, password: e.target.value })}
                    required
                    minLength={6}
                  />
                </div>
                <div className="modal-actions">
                  <button type="button" className="btn btn-outline" onClick={onClose}>Cancelar</button>
                  <button type="submit" className="btn btn-primary" disabled={createLoading}>
                    {createLoading ? 'Creando...' : 'Crear Admin'}
                  </button>
                </div>
              </form>
            )}

            {tab === 'promover' && (
              <form onSubmit={handlePromote}>
                <p className="modal-desc">Ingresá el email de un usuario ya registrado para darle permisos de administrador.</p>
                {promoteError && <div className="alert alert-error">{promoteError}</div>}
                <div className="form-group">
                  <label className="form-label">Email del usuario *</label>
                  <input
                    className="form-input"
                    type="email"
                    value={promoteEmail}
                    onChange={(e) => setPromoteEmail(e.target.value)}
                    placeholder="usuario@ejemplo.com"
                    required
                  />
                </div>
                <div className="modal-actions">
                  <button type="button" className="btn btn-outline" onClick={onClose}>Cancelar</button>
                  <button type="submit" className="btn btn-primary" disabled={promoteLoading}>
                    {promoteLoading ? 'Promoviendo...' : 'Dar permisos de Admin'}
                  </button>
                </div>
              </form>
            )}
          </>
        )}
      </div>
    </div>
  )
}

export default CreateAdminModal
