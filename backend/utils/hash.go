package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPassword retorna el hash SHA-256 en formato hexadecimal de la contraseña dada.
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// CheckPassword compara una contraseña en texto plano contra su hash almacenado.
func CheckPassword(password, hash string) bool {
	return HashPassword(password) == hash
}
