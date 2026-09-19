package folder

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/message"
	"dolsh/ui/screen"
)

func (m model) Update(msg tea.Msg) (screen.Model, tea.Cmd, error) {
	switch msg := msg.(type) {
	case message.FoldersMsg:
		return m.setFolders(msg.Folders), nil, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			return m.moveUp(), nil, nil
		case "down":
			return m.moveDown(), nil, nil
		case "enter":
			folder, ok := m.selectedFolder()
			if !ok {
				return m, nil, nil
			}
			return m, message.SubmitSelectFolder(folder.ID), nil
		}
	}

	return m, nil, nil
}
