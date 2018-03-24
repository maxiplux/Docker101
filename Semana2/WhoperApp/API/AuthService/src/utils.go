package main

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

/*
 * Cryptographically secure pseudorandom number generator.
 */
func GenerateRandomString(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	letters := string(SignKey)
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}

/*
 * HashPassword return a hash from given password
 */
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
 * END OF FILE!
 */
