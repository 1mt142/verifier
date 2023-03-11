package services

import (
	"crypto/tls"
	"encoding/base32"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"strings"
)

func GenerateRandomOTP() (string, error) {
	// Generate 4 random bytes
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// Convert the bytes to a 5-digit number using base32 encoding
	otp := base32.StdEncoding.EncodeToString(randomBytes)
	otp = otp[:5]

	return otp, nil
}

func GenerateOTP() string {
	// Generate a random 6-digit number
	otp := rand.Intn(900000) + 100000
	// Convert the number to a string and return it
	return strconv.Itoa(otp)
}

// SendOTPViaEmail A helper function to send an OTP via email
func SendOTPViaEmail(email, otp string, subject string) error {
	from := os.Getenv("FROM_EMAIL")
	password := os.Getenv("FROM_EMAIL_PASSWORD")
	to := email

	fmt.Println("To Email", to)
	fmt.Println("OTP is ", otp)

	// Compose the email message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		"Your OTP is " + otp

	// Connect to the SMTP server
	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth("", from, password, smtpServer)
	conn, err := smtp.Dial(smtpServer + ":587")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Authenticate with the SMTP server
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	if err = conn.StartTLS(config); err != nil {
		return err
	}
	if err = conn.Auth(auth); err != nil {
		return err
	}

	// Send the email message
	if err = conn.Mail(from); err != nil {
		return err
	}
	if err = conn.Rcpt(to); err != nil {
		return err
	}
	w, err := conn.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(strings.ReplaceAll(msg, "\n", "\r\n")))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}
