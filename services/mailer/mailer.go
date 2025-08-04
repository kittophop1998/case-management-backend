package mailer

import (
	"case-management/model"
	"crypto/tls"
	"net"
	"net/smtp"
	"strings"
)

type Mailer struct {
	Server      string
	Addr        string
	SenderEmail string
	Password    string
}

type Message struct {
	To      []string
	Subject string
	Body    string
}

type Email interface {
	SendSMTPMessage(message *Message) error
}

func NewMailer(server, addr, senderEmail, password string) Email {
	return &Mailer{
		Server:      server,
		Addr:        addr,
		SenderEmail: senderEmail,
		Password:    password,
	}
}

func (m *Mailer) SendSMTPMessage(message *Message) error {
	conn, err := net.Dial("tcp", m.Addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	c, err := smtp.NewClient(conn, m.Server)
	if err != nil {
		panic(err)
	}

	tlsconfig := &tls.Config{
		ServerName: m.Server,
		MinVersion: tls.VersionTLS13,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
		panic(err)
	}

	auth := LoginAuth(m.SenderEmail, m.Password)
	if err := c.Auth(auth); err != nil {
		panic(err)
	}

	// Compose email message
	msg := []byte("To: " + strings.Join(message.To, ",") + "\r\n" +
		"Content-Type: text/html; charset=UTF-8" + "\r\n" +
		"Subject: " + message.Subject + "\r\n" +
		"\r\n" +
		message.Body + "\r\n")

	err = smtp.SendMail(m.Addr, auth, m.SenderEmail, message.To, msg)
	if err != nil {
		return err
	}

	return nil
}

func LoginAuth(username, password string) smtp.Auth {
	return &model.LoginAuth{
		Username: username,
		Password: password,
	}
}
