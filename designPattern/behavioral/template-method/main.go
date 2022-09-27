package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{
		otper: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &mail{}
	o = otp{
		otper: emailOTP,
	}
	o.genAndSendOTP(4)
}
