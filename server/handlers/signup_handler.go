package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

// CreateSalt generates a random salt for Argon2
func CreateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// HashPassword uses Argon2 to hash the password with a salt
func HashPassword(password string, salt []byte) string {
	// Argon2id parameters
	time := uint32(1)           // Number of iterations
	memory := uint32(64 * 1024) // Memory cost (64MB)
	threads := uint8(4)         // Number of threads
	keyLength := uint32(32)     // Length of the generated key

	// Hash the password with Argon2 id
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	// Encode the salt and hash together for storage
	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s.%s", saltBase64, hashBase64)
}

// TODO: Signup handler
