// Package message defines messages that screens send to the root model.
package message

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/app"
)

type SetServerURL struct {
	URL string
}

type ServerURLSet struct {
	Err error
}

func SubmitServerURL(rawURL string) tea.Cmd {
	return func() tea.Msg {
		return SetServerURL{URL: rawURL}
	}
}

type Login struct {
	Username string
	Password string
}

type LoginResult struct {
	Err error
}

func SubmitLogin(username, password string) tea.Cmd {
	return func() tea.Msg {
		return Login{Username: username, Password: password}
	}
}

type SelectFolder struct {
	ID string
}

type FolderSelected struct {
	Err error
}

func SubmitSelectFolder(id string) tea.Cmd {
	return func() tea.Msg {
		return SelectFolder{ID: id}
	}
}

type LoadFolders struct{}

func SubmitLoadFolders() tea.Cmd {
	return func() tea.Msg {
		return LoadFolders{}
	}
}

type FoldersMsg struct {
	Folders []app.Folder
	Err     error
}

type LoadAlbums struct{}

func SubmitLoadAlbums() tea.Cmd {
	return func() tea.Msg {
		return LoadAlbums{}
	}
}

type AlbumsMsg struct {
	Albums []app.Album
	Err    error
}
