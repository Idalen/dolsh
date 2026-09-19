// Package folder provides the folder selection UI for dolsh.
package folder

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/message"
)

type model struct {
	folders  []app.Folder
	selector int
}

func New() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return message.SubmitLoadFolders()
}

func (m model) setFolders(folders []app.Folder) model {
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

func (m model) selectedFolder() (app.Folder, bool) {
	if m.selector < 0 || m.selector >= len(m.folders) {
		return app.Folder{}, false
	}
	return m.folders[m.selector], true
}
