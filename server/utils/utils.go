package utils

import (
	"crypto/rand"
	"encoding/hex"
	"net/smtp"
	"os"
)

func SendMail(to []string, subject string, body string) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("FROM_EMAIL"),
		os.Getenv("FROM_EMAIL_PASSWORD"),
		os.Getenv("FROM_EMAIL_SMTP"),
	)

	message := "Subject: " + subject + "\n" + body

	return smtp.SendMail(os.Getenv("SMTP_ADDRESS"), auth, os.Getenv("FROM_EMAIL"), to, []byte(message))
}

func SendHTMLMail(to []string, subject string, htmlBody string) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("FROM_EMAIL"),
		os.Getenv("FROM_EMAIL_PASSWORD"),
		os.Getenv("FROM_EMAIL_SMTP"),
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	message := "Subject: " + subject + "\n" + headers + "\n\n" + htmlBody

	return smtp.SendMail(os.Getenv("SMTP_ADDRESS"), auth, os.Getenv("FROM_EMAIL"), to, []byte(message))
}

func GenerateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
