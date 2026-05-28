package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestRegister_Success verifies that a new user can register with valid data.
func TestRegister_Success(t *testing.T) {
	// TODO: set up in-memory DB or mock UserDAO
	// TODO: create AuthService with mock
	// TODO: call service.Register with valid RegisterInput
	// TODO: assert no error and returned user has correct email and role
	assert.True(t, true, "placeholder — implement test body")
}

// TestRegister_DuplicateEmail verifies that registering with an existing email returns an error.
func TestRegister_DuplicateEmail(t *testing.T) {
	// TODO: seed DB with existing user email
	// TODO: call service.Register with same email
	// TODO: assert error is returned
	assert.True(t, true, "placeholder — implement test body")
}

// TestLogin_Success verifies that a user with valid credentials receives a JWT.
func TestLogin_Success(t *testing.T) {
	// TODO: seed user with known email and hashed password
	// TODO: call service.Login with correct credentials
	// TODO: assert no error and token is non-empty
	assert.True(t, true, "placeholder — implement test body")
}

// TestLogin_WrongPassword verifies that wrong credentials return an error.
func TestLogin_WrongPassword(t *testing.T) {
	// TODO: seed user, call service.Login with wrong password
	// TODO: assert error is returned and token is empty
	assert.True(t, true, "placeholder — implement test body")
}
