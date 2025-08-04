package model

import (
	"errors"
	"net/smtp"
)

type Email struct {
	To              []string
	From            []string
	Subject         string
	Body            string
	CarbonCopy      string `json:"cc"`
	BlindCarbonCopy string `json:"bcc"`
}

type EmailTemplate struct {
	Template        string
	Subject         string
	Body            string
	UpdatedBy       string `gorm:"-" json:"-"`
	RevisionHistory string `gorm:"-" json:"-"`
}

type LoginAuth struct {
	Username string
	Password string
}

func (a *LoginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.Username), nil
}

func (a *LoginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.Username), nil
		case "Password:":
			return []byte(a.Password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

func (EmailTemplate) TableName() string {
	return "email_template"
}
