package main

import (
	"bytes"
	"embed"
	"html/template"
	"os"
	"strconv"
	"time"

	"github.com/vanng822/go-premailer/premailer"
	mail "github.com/xhit/go-simple-mail/v2"
)

type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
}

type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
}

//go:embed templates/*.gohtml
var embeddedTemplates embed.FS

func NewMail() (Mail, error) {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		return Mail{}, err
	}

	m := Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        port,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("MAIL_ENCRYPTION"),
		FromName:    os.Getenv("MAIL_FROM_NAME"),
		FromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
	}

	return m, nil
}

func (m *Mail) SendSMTPMessage(message Message) error {
	if message.From == "" {
		message.From = m.FromAddress
	}

	if message.FromName == "" {
		message.FromName = m.FromName
	}

	message.DataMap = map[string]any{
		"message": message.Data,
	}

	htmlMessage, err := m.buildHTMLMessage(message)
	if err != nil {
		return err
	}

	plainTextMessage, err := m.buildPlainTextMessage(message)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return err
	}

	email := mail.NewMSG()
	email.SetFrom(message.From).
		AddTo(message.To).
		SetSubject(message.Subject).
		SetBody(mail.TextPlain, plainTextMessage).
		AddAlternative(mail.TextHTML, htmlMessage)

	if len(message.Attachments) > 0 {
		for _, attachment := range message.Attachments {
			email.AddAttachment(attachment)
		}
	}

	if err := email.Send(smtpClient); err != nil {
		return err
	}

	return nil
}

func (m *Mail) buildHTMLMessage(message Message) (string, error) {
	t, err := template.ParseFS(embeddedTemplates, "templates/mail.html.gohtml")
	if err != nil {
		return "", err
	}

	var tmpl bytes.Buffer
	if err = t.ExecuteTemplate(&tmpl, "body", message.DataMap); err != nil {
		return "", err
	}

	htmlMessage, err := m.inlineCSS(tmpl.String())
	if err != nil {
		return "", err
	}

	return htmlMessage, nil
}

func (m *Mail) inlineCSS(formattedMessage string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBangImportant: true,
	}

	p, err := premailer.NewPremailerFromString(formattedMessage, &options)
	if err != nil {
		return "", err
	}

	html, err := p.Transform()
	if err != nil {
		return "", err
	}

	return html, nil
}

func (m *Mail) buildPlainTextMessage(message Message) (string, error) {
	t, err := template.ParseFS(embeddedTemplates, "templates/mail.plain.gohtml")
	if err != nil {
		return "", err
	}

	var tmpl bytes.Buffer
	if err := t.ExecuteTemplate(&tmpl, "body", message.DataMap); err != nil {
		return "", err
	}

	plainTextMessage := tmpl.String()
	return plainTextMessage, nil
}

func (m *Mail) getEncryption(encryption string) mail.Encryption {
	switch encryption {
	case "tls":
		return mail.EncryptionSTARTTLS
	case "ssl":
		return mail.EncryptionSSLTLS
	case "none", "":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
