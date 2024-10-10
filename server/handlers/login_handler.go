package handlers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

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

// TODO: Login Handler
