// Package library provides the media library browsing UI for dolsh.
package library

import (
	"context"

	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/component"
)

type model struct {
	app      *app.App
	ctx      context.Context
	cancel   context.CancelFunc
	err      error
	size     component.Size
	folders  []app.Library
	selector int
}

func New(app *app.App, size component.Size) model {
	ctx, cancel := context.WithCancel(context.Background())
	return model{
		app:    app,
		ctx:    ctx,
		cancel: cancel,
		size:   size,
	}
}

func (m model) Init() tea.Cmd {
	return m.loadFolders
}

func (m model) Cancel() {
	m.cancel()
}

// setFolders replaces the folder list and keeps the selection in range.
func (m model) setFolders(folders []app.Library) model {
	m.folders = folders
	if m.selector >= len(folders) {
		m.selector = 0
	}
	return m
}

func (m model) moveUp() model {
	if len(m.folders) == 0 {
		return m
	}
	m.selector = (m.selector - 1 + len(m.folders)) % len(m.folders)
	return m
}

func (m model) moveDown() model {
	if len(m.folders) == 0 {
		return m
	}
	m.selector = (m.selector + 1) % len(m.folders)
	return m
}

func (m model) selectedFolder() (app.Library, bool) {
	if m.selector < 0 || m.selector >= len(m.folders) {
		return app.Library{}, false
	}
	return m.folders[m.selector], true
}
