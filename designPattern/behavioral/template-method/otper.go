package main

type otper interface {
	getRandOtp(int) string
	saveOtpcache(string)
	getMessage(string) string
	sendNotification(string) error
}

type otp struct {
	otper
}

func (o *otp) genAndSendOTP(optLength int) error {
	otp := o.otper.getRandOtp(5)
	o.otper.saveOtpcache(otp)
	msg := o.otper.getMessage(otp)
	return o.otper.sendNotification(msg)
}
