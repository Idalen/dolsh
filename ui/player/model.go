// Package player provides the media playback UI for dolsh.
package player

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/app"
)

type model struct {
	app *app.App
}

func New(app *app.App) model {
	return model{app: app}
}

func (m model) Init() tea.Cmd {
	return nil
}
