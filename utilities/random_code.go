package utilities

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomCode() string {
	rand.Seed(time.Now().Unix())
	charSet := "1234567890"
	var output strings.Builder
	length := 6

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}

func GenerateReferalCode() string {
	rand.Seed(time.Now().Unix())
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var output strings.Builder
	length := 9

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}

	return output.String()
}
