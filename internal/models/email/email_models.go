package emailModels

type MailData struct {
	To       string
	From     string
	Subject  string
	Content  string
	Template string
	Data     interface{}
}

type VerificationEmailData struct {
	Subject string
	Name    string
	URL     string
}
