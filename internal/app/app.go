package app

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/monban/gutt/internal/email"
	"github.com/muesli/reflow/wordwrap"
)

type App struct {
	ml       mailList
	viewport viewport.Model
	emails   email.Emails
	index    int
	provider email.Provider
	style    lipgloss.Style
}

func (a App) Init() tea.Cmd {
	return func() tea.Msg {
		return a.provider.GetMail()
	}
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// TODO: why do we have to subtract 2 here?
		a.style = a.style.Height(msg.Height - 2)
		a.style = a.style.Width(msg.Width - 2)
		a = a.Reformat()
	case email.Emails:
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
				a = a.Reformat()
				a.viewport.SetContent(wordwrap.String(a.emails[a.ml.index].Body, a.viewport.Width))
			}
		}
	}

	return a, cmd
}

func (a App) Reformat() App {
	usableWidth := float64(a.style.GetWidth())
	a.ml.list.SetWidth(int(usableWidth * 0.4))
	a.ml.list.SetHeight(a.style.GetHeight())
	a.viewport.Width = int(usableWidth * 0.6)
	a.viewport.Height = a.style.GetHeight()
	return a
}

func (a App) View() string {
	vp := a.viewport.View()
	l := a.ml.View()

	return a.style.Render(lipgloss.JoinHorizontal(lipgloss.Top, l, vp))
}

func New(p email.Provider) App {
	return App{
		ml:       NewMailList(email.Emails{}),
		viewport: viewport.New(0, 0),
		provider: p,
		style: lipgloss.NewStyle().
			Border(lipgloss.DoubleBorder()).
			Align(lipgloss.Left),
	}
}
