package email

import "time"

type Provider interface {
	GetMail() Emails
}

type Email struct {
	From    string
	Subject string
	Body    string
	Time    time.Time
}

type Emails []Email

func (e Email) FilterValue() string {
	return e.Subject
}

func (e Email) Title() string {
	return e.Subject
}

func (e Email) Description() string {
	return e.Time.Local().Format(time.DateTime)
}
