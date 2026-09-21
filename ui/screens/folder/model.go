// Package folder provides the folder selection UI for dolsh.
package folder

import (
	"context"

	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/message"
	"dolsh/ui/screens"
)

type folderService interface {
	ListFolders(ctx context.Context) ([]app.Folder, error)
	SelectFolder(id string) error
}

type model struct {
	folders  []app.Folder
	selector int
	size     screens.Size
	err      error

	svc folderService
	ctx context.Context
}

func New(svc folderService, ctx context.Context) model {
	return model{svc: svc, ctx: ctx}
}

func (m model) Init() tea.Cmd {
	return m.load()
}

func (m model) load() tea.Cmd {
	return func() tea.Msg {
		folders, err := m.svc.ListFolders(m.ctx)
		if err != nil {
			return message.Error{Err: err}
		}
		return message.Folders{Folders: folders}
	}
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
