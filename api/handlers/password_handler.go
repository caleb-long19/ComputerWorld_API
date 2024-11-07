package handlers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
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

func VerifyPassword(password, storedHash string) bool {
	// Split the stored hash into salt and actual hash
	parts := strings.Split(storedHash, ".")
	if len(parts) != 2 {
		return false
	}

	// Decode the base64 salt and hash
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return false
	}
	storedHashBytes, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Error decoding hash:", err)
		return false
	}

	// Hash the input password with the same salt
	hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	// Compare the stored hash with the newly hashed password
	return bytes.Equal(hashedPassword, storedHashBytes)
}
