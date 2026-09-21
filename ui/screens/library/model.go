// Package library provides the media library browsing UI for dolsh.
package library

import (
	"context"

	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/message"
	"dolsh/ui/screens"
)

type libraryService interface {
	Albums(ctx context.Context) ([]app.Album, error)
	Tracks(ctx context.Context, albumID string) ([]app.Track, error)
}

type model struct {
	albums   []app.Album
	selector int
	tracks   map[string][]app.Track
	size     screens.Size
	err      error

	svc libraryService
	ctx context.Context
}

func New(svc libraryService, ctx context.Context) model {
	return model{
		svc:    svc,
		ctx:    ctx,
		tracks: make(map[string][]app.Track),
	}
}

func (m model) Init() tea.Cmd {
	return m.load()
}

func (m model) load() tea.Cmd {
	return func() tea.Msg {
		albums, err := m.svc.Albums(m.ctx)
		if err != nil {
			return message.Error{Err: err}
		}
		return message.Albums{Albums: albums}
	}
}

func (m model) loadTracks(albumID string) tea.Cmd {
	return func() tea.Msg {
		tracks, err := m.svc.Tracks(m.ctx, albumID)
		if err != nil {
			return message.Error{Err: err}
		}
		return message.Tracks{AlbumID: albumID, Tracks: tracks}
	}
}

func (m model) loadSelected() (model, tea.Cmd) {
	album, ok := m.selectedAlbum()
	if !ok {
		return m, nil
	}

	if _, ok := m.tracks[album.ID]; ok {
		return m, nil
	}

	return m, m.loadTracks(album.ID)
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
