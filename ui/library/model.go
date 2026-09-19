// Package library provides the media library browsing UI for dolsh.
package library

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/message"
)

type model struct {
	albums   []app.Album
	selector int
}

func New() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return message.SubmitLoadAlbums()
}

func (m model) setAlbums(albums []app.Album) model {
	m.albums = albums
	if m.selector >= len(albums) {
		m.selector = 0
	}
	return m
}

func (m model) moveUp() model {
	if len(m.albums) == 0 {
		return m
	}
	m.selector = (m.selector - 1 + len(m.albums)) % len(m.albums)
	return m
}

func (m model) moveDown() model {
	if len(m.albums) == 0 {
		return m
	}
	m.selector = (m.selector + 1) % len(m.albums)
	return m
}

func (m model) selectedAlbum() (app.Album, bool) {
	if m.selector < 0 || m.selector >= len(m.albums) {
		return app.Album{}, false
	}
	return m.albums[m.selector], true
}
