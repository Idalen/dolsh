package folder

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

	case message.Folders:
		return m.setFolders(msg.Folders), nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			return m.moveUp(), nil
		case "down":
			return m.moveDown(), nil
		case "enter":
			folder, ok := m.selectedFolder()
			if !ok {
				return m, nil
			}
			return m, m.selectFolder(folder.ID)
		}
	}

	return m, nil
}

func (m model) selectFolder(id string) tea.Cmd {
	return func() tea.Msg {
		if err := m.svc.SelectFolder(id); err != nil {
			return message.Error{Err: err}
		}
		return message.Next{}
	}
}
