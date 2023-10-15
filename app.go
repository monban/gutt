package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
)

var (
	mlStyle = lipgloss.NewStyle().Border(lipgloss.DoubleBorder()).Align(lipgloss.Left)
	vpStyle = mlStyle.Copy().Align(lipgloss.Right)
)

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
}

func (a app) Init() tea.Cmd {
	return nil
}

func (a app) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		a.geo.height = msg.Height - 2
		a.geo.width = msg.Width
		a = a.Reformat()
	}
	var cmd tea.Cmd
	a.ml, cmd = a.ml.Update(msg)
	a.viewport.SetContent(wordwrap.String(a.emails[a.ml.index].body, a.viewport.Width-4))
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

func NewApp(emails []email) app {
	return app{
		emails:   emails,
		ml:       NewMailList(emails),
		viewport: viewport.New(0, 0),
	}
}
