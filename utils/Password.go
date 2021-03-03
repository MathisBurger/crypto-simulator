package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
)

// hashes the password
func HashPassword(pwd string) string {
	salt := generateSalt()
	hash := hashPassword(pwd, salt)
	for i := 0; i < 6; i++ {
		hash = hashPassword(hash, salt)
	}
	return hash

}

// generates salt
func generateSalt() []byte {
	var salt = make([]byte, 16)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err.Error())
	}
	return salt
}

// hashes with salt
func hashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var base64EncodedPasswordHash = hex.EncodeToString(hashedPasswordBytes)
	var encodedSalt = hex.EncodeToString(salt)
	return encodedSalt + "$" + base64EncodedPasswordHash
}

// checks if passwords are the same
func DoPasswordsMatch(hashedPassword, currPassword string,
	salt string) bool {
	slt, _ := hex.DecodeString(salt)
	currPasswordHash := hashPassword(currPassword, slt)
	for i := 0; i < 6; i++ {
		currPasswordHash = hashPassword(currPasswordHash, slt)
	}
	return hashedPassword == currPasswordHash
}
