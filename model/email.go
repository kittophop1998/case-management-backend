package model

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

func (EmailTemplate) TableName() string {
	return "email_template"
}
