// Package login provides the UI for jellyfin login
package login

import (
	"context"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"dolsh/ui/component"
	"dolsh/ui/screens"
)

type loginService interface {
	Login(ctx context.Context, username, password string) error
}

type model struct {
	input      [2]component.Input
	selector   int
	submitting bool
	spinner    spinner.Model
	size       screens.Size
	err        error

	svc loginService
	ctx context.Context
}

func New(svc loginService, ctx context.Context) model {
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
		svc:     svc,
		ctx:     ctx,
	}
}

func (m model) Init() tea.Cmd {
	return m.input[m.selector].Focus()
}
