package template

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
}

type Otp struct {
	Iotp IOtp
}

func (o *Otp) GenAndSendOTP(otpLength int) error {
	opt := o.Iotp.GenRandomOTP(otpLength)
	o.Iotp.SaveOTPCache(opt)
	message := o.Iotp.GetMessage(opt)
	err := o.Iotp.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
