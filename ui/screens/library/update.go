package library

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/message"
	"dolsh/ui/screens"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = screens.FromWindowSize(msg)
		return m, nil

	case message.Error:
		m.err = msg.Err
		return m, nil

	case message.Albums:
		return m.setAlbums(msg.Albums).loadSelected()

	case message.Tracks:
		m.tracks[msg.AlbumID] = msg.Tracks
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			return m.moveUp().loadSelected()
		case "down":
			return m.moveDown().loadSelected()
		}
	}

	return m, nil
}
