package emailModels

type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
	Data     interface{}
}

type VerificationCodeEmailData struct {
	Subject string
	Name    string
	Code    string
	URL     string
}
