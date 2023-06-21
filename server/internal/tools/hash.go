package tools

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

// HashString func is used to create hash based on
// string and const default is 14
func HashString(value string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// CheckPasswordHash func is used to compare to hashes
func CheckPasswordHash(value, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(value))
	return err == nil
}

// Base64Encode метод кодирующий данные в формате base64
func Base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// Base64Decode метод декодирующий данные в формате base64
func Base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
