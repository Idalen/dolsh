// Package setup provides the UI responsible for jellyfin server URL definition
package setup

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/component"
)

type model struct {
	app   *app.App
	input component.Input
	size  component.Size
	err   error
}

func New(app *app.App) model {
	return model{
		app:   app,
		input: component.NewInput("http(s)://your-jellyfin-server.com"),
	}
}

func (m model) Init() tea.Cmd {
	return m.input.Focus()
}
