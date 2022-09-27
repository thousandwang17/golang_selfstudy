package main

import "fmt"

type mail struct {
	otp otper
}

func Newmail() otper {
	return &mail{}
}

func (s *mail) getRandOtp(len int) string {
	randOtp := "1234"
	fmt.Println("mail generate opt", randOtp)
	return randOtp
}

func (s *mail) saveOtpcache(otp string) {
	fmt.Println("mail OTP saved to CACHE")
}

func (s *mail) getMessage(otp string) string {
	return "mail OTP for login is " + otp
}

func (s *mail) sendNotification(message string) error {
	fmt.Printf("mail: sending sms: %s\n", message)
	return nil
}
