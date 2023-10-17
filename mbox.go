package main

import (
	"errors"
	"io"
	"net/mail"
	"os"
	"time"

	"github.com/charmbracelet/log"
	mbox "github.com/emersion/go-mbox"
)

type MboxProvider struct {
	file string
}

func (mbp MboxProvider) GetMail() emails {
	var emails []email
	file, err := os.Open(mbp.file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := mbox.NewReader(file)
	for {
		var email email
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
		email.from = e.Header.Get("From")
		email.subject = e.Header.Get("Subject")
		body, err := io.ReadAll(e.Body)
		if err != nil {
			log.Fatal(err)
		}
		email.body = string(body)
		emails = append(emails, email)
	}
	time.Sleep(time.Second) //HACK: this is for ui testing
	return emails
}
