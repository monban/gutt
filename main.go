package main

import (
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
)

func main() {
	prov := MboxProvider{getMailfile()}
	var a app = NewApp(prov)

	p := tea.NewProgram(a, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
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
