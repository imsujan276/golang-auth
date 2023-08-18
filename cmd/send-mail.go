package main

import (
	"bytes"
	"crypto/tls"
	"log"
	"os"
	"path/filepath"
	"pomo/internal/config"
	emailModels "pomo/internal/models/email"
	emailTemplates "pomo/templates-email"
	"text/template"

	"github.com/k3a/html2text"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func listenForMail() {
	// start anonymous go routine
	go func() {
		for {
			msg := <-config.Config.MailChannel
			sendEmail(msg)
		}
	}()
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

func sendEmail(m emailModels.MailData) {
	config := config.Config

	// Sender data.
	from := config.EmailFrom
	smtpPass := config.SMTPPass
	smtpUser := config.SMTPUser
	to := m.To
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort

	var body bytes.Buffer

	template, err := ParseTemplateDir(emailTemplates.Path)
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, m.Template, &m.Data)

	mail := gomail.NewMessage()

	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", m.Subject)
	mail.SetBody("text/html", body.String())
	mail.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(mail); err != nil {
		logrus.Fatal("Could not send email: ", err)
	}
	logrus.Info("Email sent")

}

// func sendEmail(m emailModels.MailData) {
// 	server := mail.NewSMTPClient()
// 	server.Host = config.Config.SMTPHost
// 	server.Port = config.Config.SMTPPort
// 	server.Username = config.Config.SMTPUser
// 	server.Password = config.Config.SMTPPass
// 	server.Encryption = mail.EncryptionSTARTTLS
// 	server.KeepAlive = false
// 	server.ConnectTimeout = 10 * time.Second
// 	server.SendTimeout = 15 * time.Second

// 	client, err := server.Connect()
// 	if err != nil {
// 		config.Config.ErrorLog.Println(err)
// 	}

// 	email := mail.NewMSG()
// 	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

// 	if m.Template == "" {
// 		email.SetBody(mail.TextHTML, m.Content)
// 	} else {
// 		data, err := ioutil.ReadFile(fmt.Sprintf("./templates-email/%s", m.Template))
// 		if err != nil {
// 			config.Config.ErrorLog.Println(err)
// 		}
// 		mailtemplate := string(data)
// 		msgToSend := strings.Replace(mailtemplate, "[%body%]", m.Content, 1)
// 		email.SetBody(mail.TextHTML, msgToSend)
// 	}

// 	err = email.Send(client)
// 	if err != nil {
// 		config.Config.ErrorLog.Println(err)
// 	} else {
// 		config.Config.InfoLog.Println("Email sent")
// 	}
// }
