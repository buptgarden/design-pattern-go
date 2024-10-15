package template

import (
	"fmt"
)

type Email struct {
	Otp
}

func (e *Email) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random opt %s\n", randomOTP)
	return randomOTP
}

func (e *Email) SaveOTPCache(opt string) {
	fmt.Printf("EMAIL: svaing otp: %s to cache\n", opt)
}

func (e *Email) GetMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (e *Email) SendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
