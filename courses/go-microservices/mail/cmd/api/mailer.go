package main

import (
	"bytes"
	"html/template"

	"github.com/vanng822/go-premailer/premailer"
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

	formattedMessage, err := m.buildHTMLMessage(message)
}

func (m *Mail) buildHTMLMessage(message Message) (string, error) {
	templateToRender := "./templates/mail.gohtml"
	t, err := template.New("email-html").ParseFiles(templateToRender)
	if err != nil {
		return "", err
	}

	var tmpl bytes.Buffer
	if err := t.ExecuteTemplate(&tmpl, "body", message.DataMap); err != nil {
		return "", err
	}

	formattedMessage := tmpl.String()
	formattedMessage, err = m.inlineCSS(formattedMessage)
	if err != nil {
		return "", err
	}

	return formattedMessage, nil
}

func (m *Mail) inlineCSS(formattedMessage string) (string, error) {
	options := premailer.Options{
		RemoveClasses:     false,
		CssToAttributes:   false,
		KeepBandImportant: true,
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
