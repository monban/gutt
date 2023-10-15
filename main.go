package main

import (
	"errors"
	"io"
	"net/mail"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	mbox "github.com/emersion/go-mbox"
)

func main() {
	emails := readEmails()
	var a app = NewApp(emails)

	p := tea.NewProgram(a, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

func readEmails() []email {
	var emails []email
	file, err := os.Open(getMailfile())
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := mbox.NewReader(file)
	for {
		var email email
		msg, err := reader.NextMessage()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			panic(err)
		}
		e, err := mail.ReadMessage(msg)
		if err != nil {
			panic(err)
		}
		email.from = e.Header.Get("From")
		email.subject = e.Header.Get("Subject")
		body, err := io.ReadAll(e.Body)
		if err != nil {
			panic(err)
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
		panic("can't find mailfile")
	}
	return mailfile
}
