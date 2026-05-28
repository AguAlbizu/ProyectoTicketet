package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword returns the SHA-256 hex digest of the given password.
func HashPassword(password string) string {
	// TODO: compute sha256.Sum256([]byte(password))
	// TODO: return hex.EncodeToString(hash[:])
	_ = sha256.New()
	_ = hex.EncodeToString
	return ""
}

// CheckPassword compares a plain-text password against its stored hash.
func CheckPassword(password, hash string) bool {
	// TODO: hash the password and compare with the stored hash
	return false
}
