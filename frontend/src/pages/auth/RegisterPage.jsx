import { useAuth } from '../../hooks/useAuth'
import { useNavigate } from 'react-router-dom'

// Register page — creates a new client account.
function RegisterPage() {
  const { register } = useAuth()
  const navigate = useNavigate()

  // TODO: manage form state: name, email, password, confirmPassword
  // TODO: validate that passwords match before submitting
  // TODO: call register(name, email, password)
  // TODO: on success redirect to /
  // TODO: show error messages (duplicate email, validation failures)

  return (
    <div>
      {/* TODO: implement register form UI */}
    </div>
  )
}

export default RegisterPage
