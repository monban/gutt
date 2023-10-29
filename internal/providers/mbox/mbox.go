package mbox

import (
	"errors"
	"io"
	"net/mail"
	"os"
	"time"

	"github.com/charmbracelet/log"
	mbox "github.com/emersion/go-mbox"
	"github.com/monban/gutt/internal/email"
)

type MboxProvider struct {
	file string
}

func New(filename string) email.Provider {
	return MboxProvider{filename}
}

func (mbp MboxProvider) GetMail() email.Emails {
	var emails email.Emails
	file, err := os.Open(mbp.file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := mbox.NewReader(file)
	for {
		var email email.Email
		msg, err := reader.NextMessage()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		e, err := mail.ReadMessage(msg)
		if err != nil {
			log.Fatal(err)
		}
		email.From = e.Header.Get("From")
		email.Subject = e.Header.Get("Subject")
		if email.Time, err = e.Header.Date(); err != nil {
			email.Time = time.Now()
		}
		body, err := io.ReadAll(e.Body)
		if err != nil {
			log.Fatal(err)
		}
		email.Body = string(body)
		emails = append(emails, email)
	}
	return emails
}
