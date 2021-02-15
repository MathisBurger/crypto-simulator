package utils

import (
	"math/rand"
	"strings"
	"time"
)

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
