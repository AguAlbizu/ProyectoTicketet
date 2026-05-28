import { useAuth } from '../../hooks/useAuth'
import { useNavigate } from 'react-router-dom'

// Login page — authenticates an existing user.
function LoginPage() {
  const { login } = useAuth()
  const navigate = useNavigate()

  // TODO: manage form state: email, password
  // TODO: on submit call login(email, password)
  // TODO: on success redirect based on user.role:
  //         admin  → /admin
  //         cliente → /
  // TODO: show error message on failed login

  return (
    <div>
      {/* TODO: implement login form UI */}
    </div>
  )
}

export default LoginPage
