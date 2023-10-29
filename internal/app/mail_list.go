package app

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/monban/gutt/internal/email"
)

type mailList struct {
	list   list.Model
	emails email.Emails
	index  int
}

type indexUpdated int

func (ml mailList) Init() tea.Cmd {
	return nil
}

func (ml mailList) Update(msg tea.Msg) (mailList, tea.Cmd) {
	var cmd tea.Cmd
	ml.list, cmd = ml.list.Update(msg)
	ml.index = ml.list.Index()
	return ml, cmd
}

func (ml mailList) View() string {
	return ml.list.View()
}

func NewMailList(emails email.Emails) mailList {
	var ml mailList
	ml.emails = emails
	var items []list.Item = make([]list.Item, len(ml.emails))
	for i, j := range ml.emails {
		items[i] = j
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	ml.list = l
	return ml
}
