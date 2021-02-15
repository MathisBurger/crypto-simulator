package utils

import (
	"crypto/sha512"
	"encoding/base64"
	"math/rand"
)

func HashPassword(pwd string) string {
	salt := generateSalt()
	hash := hashPassword(pwd, salt)
	return hash

}

func generateSalt() []byte {
	var salt = make([]byte, 16)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err.Error())
	}
	return salt
}

func hashPassword(password string, salt []byte) string {
	var passwordBytes = []byte(password)
	var sha512Hasher = sha512.New()
	passwordBytes = append(passwordBytes, salt...)
	sha512Hasher.Write(passwordBytes)
	var hashedPasswordBytes = sha512Hasher.Sum(nil)
	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	var encodedSalt = base64.URLEncoding.EncodeToString(salt)

	return encodedSalt + "$" + base64EncodedPasswordHash
}

func doPasswordsMatch(hashedPassword, currPassword string,
	salt []byte) bool {
	var currPasswordHash = hashPassword(currPassword, salt)

	return hashedPassword == currPasswordHash
}
