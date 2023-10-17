package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

var (
	mlStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Align(lipgloss.Left)
	vpStyle = mlStyle.Copy().Align(lipgloss.Right)
)

type Provider interface {
	GetMail() emails
}

type geometry struct {
	width  int
	height int
}

type app struct {
	ml       mailList
	viewport viewport.Model
	emails   []email
	index    int
	geo      geometry
	provider Provider
}

type emails []email

func (a app) Init() tea.Cmd {
	return func() tea.Msg {
		return a.provider.GetMail()
	}
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.geo.height = msg.Height - 2
		a.geo.width = msg.Width
		a = a.Reformat()
	case emails:
		a.emails = msg
		a.ml = NewMailList(a.emails)
		a = a.Reformat()
	default:
		i := a.ml.index
		a.ml, cmd = a.ml.Update(msg)

		// TODO: checking if a variable has changed seems
		// the wrong way to go about this
		if a.ml.index != i {
			if len(a.emails) >= a.ml.index {
				a.viewport.SetContent(wordwrap.String(a.emails[a.ml.index].body, a.viewport.Width-4))
			}
		}
	}

	return a, cmd
}

func (a app) Reformat() app {
	a.ml.list.SetWidth(int(float64(a.geo.width) * 0.4))
	a.ml.list.SetHeight(a.geo.height)
	a.viewport.Width = int(float64(a.geo.width) * 0.6)
	a.viewport.Height = a.geo.height
	return a
}

func (a app) View() string {
	vp := vpStyle.Render(a.viewport.View())
	l := vpStyle.Render(a.ml.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, l, vp)
}

func NewApp(p Provider) app {
	return app{
		ml:       NewMailList([]email{}),
		viewport: viewport.New(0, 0),
		provider: p,
	}
}
