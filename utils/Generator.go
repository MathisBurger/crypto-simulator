package utils

import (
	"encoding/base64"
	"math/rand"
	"strings"
	"time"
)

// UUID generator
func GenerateUUID() string {
	rand.Seed(time.Now().Unix())
	charSet := "abcdefghijklmnopqrstuvwxyz"
	var output = strings.Builder{}
	partLength := 4
	numParts := 6
	for i := 0; i < numParts; i++ {
		if i != 0 {
			output.WriteString("-")
		}
		for j := 0; j < partLength; j++ {
			random := rand.Intn(len(charSet))
			randomChar := charSet[random]
			output.WriteString(string(randomChar))
		}
	}
	return output.String()
}

func GenerateToken() string {
	rand.Seed(time.Now().Unix())
	charSet := "abcdefghijklmnopqrstuvwxyz"
	var output = strings.Builder{}
	for i := 0; i < 128; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

func ByteArray(lngth int) ([]byte, error) {
	arr := make([]byte, lngth)
	_, err := rand.Read(arr)
	return arr, err
}

func Base64(lngth int) string {
	str, _ := ByteArray(lngth)
	return base64.StdEncoding.EncodeToString(str)
}
