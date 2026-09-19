package library

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/message"
	"dolsh/ui/screen"
)

func (m model) Update(msg tea.Msg) (screen.Model, tea.Cmd, error) {
	switch msg := msg.(type) {
	case message.AlbumsMsg:
		return m.setAlbums(msg.Albums), nil, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			return m.moveUp(), nil, nil
		case "down":
			return m.moveDown(), nil, nil
		}
	}

	return m, nil, nil
}
