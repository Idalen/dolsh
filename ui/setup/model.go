// Package setup provides the UI responsible for jellyfin server URL definition
package setup

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/component"
)

type model struct {
	input component.Input
}

func New() model {
	return model{
		input: component.NewInput("http(s)://your-jellyfin-server.com"),
	}
}

func (m model) Init() tea.Cmd {
	return m.input.Focus()
}
