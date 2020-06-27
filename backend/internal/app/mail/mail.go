package mail

import (
	"bytes"
	"html/template"
	"net/smtp"
)

// Mail ...
type Mail struct {
	Auth           smtp.Auth
	Host           string
	Port           string
	Email          string
	SignupTemplate *template.Template
	ForgotTemplate *template.Template
}

// Config ...
type Config struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	SignupTemplate string `json:"signupTemplate"`
	ForgotTemplate string `json:"forgotTemplate"`
}

// New ...
func New(config *Config) (*Mail, error) {

	confirmTempl, err := template.ParseFiles(config.SignupTemplate)
	if err != nil {
		return nil, err
	}

	forgotTempl, err := template.ParseFiles(config.ForgotTemplate)
	if err != nil {
		return nil, err
	}

	return &Mail{
		Auth: smtp.PlainAuth(
			"",
			config.Email,
			config.Password,
			config.Host,
		),
		Host:           config.Host,
		Port:           config.Port,
		Email:          config.Email,
		SignupTemplate: confirmTempl,
		ForgotTemplate: forgotTempl,
	}, nil
}

// Ping ...
func (mail *Mail) Ping() error {

	err := smtp.SendMail(
		mail.Host+mail.Port,
		mail.Auth,
		mail.Email,
		[]string{mail.Email},
		[]byte("From: Ping <"+mail.Email+"> \r\n"+
			"To:"+mail.Email+"\r\n"+
			"Subject: Mail Check\r\n"),
	)
	if err != nil {
		return err
	}
	return nil
}

// SendingURLBusiness ...
func (m *Mail) SendingConfirmURL(to string, url string) error {

	mime := "MIME-version: 1.0;\nContent-Type: text/html;\n\n"
	subject := "Subject: Регистрация на LSAIT!\n"

	buf := new(bytes.Buffer)
	if err := m.SignupTemplate.Execute(buf, struct{ URL string }{URL: url}); err != nil {
		return err
	}

	msg := []byte(subject + mime + buf.String())
	err := smtp.SendMail(
		m.Host+m.Port,
		m.Auth,
		m.Email,
		[]string{to},
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mail) SendingForgotURL(to string, url string) error {

	mime := "MIME-version: 1.0;\nContent-Type: text/html;\n\n"
	subject := "Subject: Смена пароля на LSAIT\n"

	buf := new(bytes.Buffer)
	if err := m.ForgotTemplate.Execute(buf, struct{ URL string }{URL: url}); err != nil {
		return err
	}

	msg := []byte(subject + mime + buf.String())
	err := smtp.SendMail(
		m.Host+m.Port,
		m.Auth,
		m.Email,
		[]string{to},
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}
