// Package message defines messages that screens send to the root model.
package message

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/app"
)

// ErrorMsg reports an error to the root model.
type ErrorMsg struct {
	Err error
}

type SetServerURL struct {
	URL string
}

type ServerURLSet struct{}

func SubmitServerURL(rawURL string) tea.Cmd {
	return func() tea.Msg {
		return SetServerURL{URL: rawURL}
	}
}

type Login struct {
	Username string
	Password string
}

type LoginResult struct{}

func SubmitLogin(username, password string) tea.Cmd {
	return func() tea.Msg {
		return Login{Username: username, Password: password}
	}
}

type SelectFolder struct {
	ID string
}

type FolderSelected struct{}

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
}

type LoadAlbums struct{}

func SubmitLoadAlbums() tea.Cmd {
	return func() tea.Msg {
		return LoadAlbums{}
	}
}

type AlbumsMsg struct {
	Albums []app.Album
}

type TracksMsg struct {
	Tracks []app.Track
}
