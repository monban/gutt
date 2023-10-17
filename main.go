package main

import (
	"errors"
	"io"
	"net/mail"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	mbox "github.com/emersion/go-mbox"
)

func main() {
	emails := readEmails()
	var a app = NewApp(emails)

	p := tea.NewProgram(a, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func readEmails() []email {
	var emails []email
	file, err := os.Open(getMailfile())
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
	return emails
}

func getMailfile() string {
	var mailfile string
	mailfile = os.Getenv("MAIL")
	if mailfile == "" {
		mailfile = filepath.Join("/var/spool/mail", os.Getenv("USER"))
	}
	if mailfile == "" {
		log.Fatal("can't find mailfile")
	}
	return mailfile
}
