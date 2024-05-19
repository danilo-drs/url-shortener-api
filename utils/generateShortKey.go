package utils

import (
	"crypto/rand"
	"log"
	"os"
	"strconv"
)

func GenerateShortKey() string {
	// get the pattern of the short URL from the environment variables
	pattern := os.Getenv("SHORT_URL_STRING")
	if pattern == "" {
		// default pattern for short URL
		pattern = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	// get the short URL length from the environment variables
	shortUrlLengthString := os.Getenv("SHORT_URL_LENGTH")
	// get the integer value of the short URL length
	shortUrlLength, err := strconv.Atoi(shortUrlLengthString)
	if err != nil {
		shortUrlLength = 6 // default value for short URL length
	}
	// generate the short URL
	bytes := make([]byte, shortUrlLength)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	for i, b := range bytes {
		bytes[i] = pattern[b%byte(len(pattern))]
	}
	// return the short URL as a string
	return string(bytes)
}
