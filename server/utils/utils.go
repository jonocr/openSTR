package utils

import (
	"crypto/rand"
	"encoding/hex"
	"net/smtp"
	"os"
)

func SendMail(to []string, body string) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("FROM_EMAIL"),
		os.Getenv("FROM_EMAIL_PASSWORD"),
		os.Getenv("FROM_EMAIL_SMTP"),
	)

	return smtp.SendMail(os.Getenv("SMTP_ADDRESS"), auth, os.Getenv("FROM_EMAIL"), to, []byte(body))
}

func GenerateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
