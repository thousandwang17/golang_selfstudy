package main

import "fmt"

type sms struct{}

func NewSms() otper {
	return &sms{}
}

func (s *sms) getRandOtp(len int) string {
	randOtp := "1234"
	fmt.Println("sms generate opt", randOtp)
	return randOtp
}

func (s *sms) saveOtpcache(otp string) {
	fmt.Println("sms OTP saved to CACHE")
}

func (s *sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
