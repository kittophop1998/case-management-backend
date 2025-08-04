package mailer

import (
	"net/smtp"
	"strings"
)

type MailTrap struct {
	Server      string
	Addr        string
	SenderEmail string
	Password    string
}

func NewMailtTrap(server, addr, senderEmail, password string) Email {
	return &MailTrap{
		Server:      server,
		Addr:        addr,
		SenderEmail: senderEmail,
		Password:    password,
	}
}

// func NewMailtTrap(host, port, username, password string) *Mailer {
// 	return &Mailer{
// 		Server:      host,
// 		Addr:        host + ":" + port,
// 		SenderEmail: username,
// 		Password:    password,
// 	}
// }

func (m *MailTrap) SendSMTPMessage(message *Message) error {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "john.doe@gmail.com", "extremely_secret_pass", "smtp.gmail.com")

	// Compose email message
	msg := []byte("To: " + strings.Join(message.To, ",") + "\r\n" +
		"Subject: " + message.Subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8" +
		"\r\n" +
		message.Body + "\r\n")

	err := smtp.SendMail(m.Addr, auth, m.SenderEmail, message.To, msg)
	if err != nil {
		return err
	}
	return err
}

// func (m *Mailer) SendSMTPMessage(message *Message) error {
// 	// เชื่อมต่อกับ SMTP server โดยใช้ host:port (เช่น sandbox.smtp.mailtrap.io:2525)
// 	conn, err := net.Dial("tcp", m.Addr)
// 	if err != nil {
// 		return err
// 	}
// 	defer conn.Close()

// 	// สร้าง client ใหม่ โดยใช้เฉพาะ hostname (ไม่เอา port)
// 	c, err := smtp.NewClient(conn, m.Server) // m.Server ต้องเป็น hostname ไม่มี port
// 	if err != nil {
// 		return err
// 	}
// 	defer c.Close()

// 	// สร้าง tls config โดย ServerName ต้องเป็น hostname เท่านั้น
// 	tlsconfig := &tls.Config{
// 		ServerName: m.Server, // ตัวอย่าง "sandbox.smtp.mailtrap.io" (ไม่มี port)
// 		MinVersion: tls.VersionTLS13,
// 	}

// 	if err = c.StartTLS(tlsconfig); err != nil {
// 		return err
// 	}

// 	auth := LoginAuth(m.SenderEmail, m.Password)
// 	if err := c.Auth(auth); err != nil {
// 		return err
// 	}

// 	// สร้างข้อความอีเมล
// 	msg := []byte("To: " + strings.Join(message.To, ",") + "\r\n" +
// 		"Content-Type: text/html; charset=UTF-8\r\n" +
// 		"Subject: " + message.Subject + "\r\n" +
// 		"\r\n" +
// 		message.Body + "\r\n")

// 	// ส่งอีเมล (คุณอาจใช้ c.Mail(), c.Rcpt(), c.Data() ส่งแทน smtp.SendMail)
// 	if err := smtp.SendMail(m.Addr, auth, m.SenderEmail, message.To, msg); err != nil {
// 		return err
// 	}

// 	return nil
// }
