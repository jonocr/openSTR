package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"text/template"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

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

func SendHTMLTemplateMail(to []string, subject string, emailData map[string]string, templateName string) error {
	auth := smtp.PlainAuth(
		"",
		os.Getenv("FROM_EMAIL"),
		os.Getenv("FROM_EMAIL_PASSWORD"),
		os.Getenv("FROM_EMAIL_SMTP"),
	)

	template, err := template.ParseFiles("./templates/" + templateName + ".html")

	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	//Render the template with data
	var renderedTemplate bytes.Buffer
	if err := template.Execute(&renderedTemplate, emailData); err != nil {
		log.Println(err.Error())
		return err
	}

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	message := "Subject: " + subject + "\n" + headers + "\n\n" + renderedTemplate.String()

	return smtp.SendMail(os.Getenv("SMTP_ADDRESS"), auth, os.Getenv("FROM_EMAIL"), to, []byte(message))
}

func GenerateRandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func CreateToken(username string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Minute * 2).Unix(),
		})
	fmt.Printf(Blue+"Secret: %+v \n"+Reset, secretKey)
	fmt.Printf(Yellow+"Token: %+v \n"+Reset, token)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Printf(Red+"Error from SignedString: %+v \n"+Reset, err.Error())
		return "", err
	}

	return tokenString, nil
}
