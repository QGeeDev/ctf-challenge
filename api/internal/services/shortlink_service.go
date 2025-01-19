package services

import (
	"crypto/rand"
	"encoding/base64"
)

func GetFullUrlByShortlinkSlug(slug string) string {
	//TODO: Query DB for URL
	// If null, return empty string
	return slug
}

func CreateShortlink(targetUrl string) (string, error) {
	result, err := generateBase64String(8)
	if err != nil {
		return "", err
	}
	// TODO: Write URL to DB
	return result, err
}

func generateBase64String(length int) (string, error) {
	byteLength := (length * 3) / 4

	randomBytes := make([]byte, byteLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	encoded := base64.RawURLEncoding.EncodeToString(randomBytes)
	return encoded[:length], nil
}
