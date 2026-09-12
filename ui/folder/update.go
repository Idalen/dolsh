package library

import (
	"log"

	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/component"
	"dolsh/ui/player"
	"dolsh/ui/screen"
)

type foldersMsg struct {
	folders []app.Folder
	err     error
}

func (m model) loadFolders() tea.Msg {
	folders, err := m.app.ListFolders(m.ctx)
	return foldersMsg{folders: folders, err: err}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case foldersMsg:
		if msg.err != nil {
			log.Printf("list folders failed: %v", msg.err)
			m.err = msg.err
			return m, nil
		}
		m = m.setFolders(msg.folders)
		return m, nil

	case tea.WindowSizeMsg:
		m.size = component.FromWindowSize(msg)
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Cancel()
			return m, tea.Quit
		case "up":
			return m.moveUp(), nil
		case "down":
			return m.moveDown(), nil
		case "enter":
			lib, ok := m.selectedFolder()
			if !ok {
				return m, nil
			}
			if err := m.app.SelectLibrary(lib.ID); err != nil {
				m.err = err
				return m, nil
			}
			return screen.Switch(m, player.New(m.app))
		}
	}

	return m, nil
}
