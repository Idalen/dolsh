// Package setup provides the UI responsible for jellyfin server URL definition
package setup

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/component"
	"dolsh/ui/screens"
)

type setupService interface {
	SetServerURL(rawURL string) error
}

type model struct {
	input component.Input
	size  screens.Size
	err   error
	svc   setupService
}

func New(svc setupService) model {
	return model{
		input: component.NewInput("http(s)://your-jellyfin-server.com"),
		svc:   svc,
	}
}

func (m model) Init() tea.Cmd {
	return m.input.Focus()
}
