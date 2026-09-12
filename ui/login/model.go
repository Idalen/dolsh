// Package login provides the UI for jellyfin login
package login

import (
	"context"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"dolsh/app"
	"dolsh/ui/component"
)

type model struct {
	app    *app.App
	ctx    context.Context
	cancel context.CancelFunc
	err    error
	// UI
	input      [2]component.Input
	selector   int
	size       component.Size
	submitting bool
	spinner    spinner.Model
}

func New(app *app.App, size component.Size) model {
	ctx, cancel := context.WithCancel(context.Background())
	input := [2]component.Input{
		component.NewInput("username"),
		component.NewPasswordInput("password"),
	}
	input[1].Blur()

	sp := spinner.New()
	sp.Spinner = spinner.Dot
	sp.Style = lipgloss.NewStyle().Foreground(component.Accent)

	m := model{
		app:     app,
		ctx:     ctx,
		cancel:  cancel,
		input:   input,
		spinner: sp,
	}
	return m.withSize(size)
}

func (m model) withSize(size component.Size) model {
	m.size = size
	for i := range m.input {
		m.input[i] = m.input[i].SetWidth(contentWidth(size.Width))
	}
	return m
}

func (m model) Init() tea.Cmd {
	return m.input[m.selector].Focus()
}

func (m model) Cancel() {
	m.cancel()
}
