// Package login provides the UI for jellyfin login
package login

import (
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"dolsh/ui/component"
)

type model struct {
	input      [2]component.Input
	selector   int
	submitting bool
	spinner    spinner.Model
}

func New() model {
	input := [2]component.Input{
		component.NewInput("username"),
		component.NewPasswordInput("password"),
	}
	input[1].Blur()

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(component.Accent)

	return model{
		input:   input,
		spinner: sp,
	}
}

func (m model) Init() tea.Cmd {
	return m.input[m.selector].Focus()
}
