package ui

import (
	"context"

	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/folder"
	"dolsh/ui/library"
	"dolsh/ui/login"
	"dolsh/ui/message"
	"dolsh/ui/screen"
	"dolsh/ui/setup"
)

type model struct {
	screen screen.Model

	app    *app.App
	size   screen.Size
	ctx    context.Context
	cancel context.CancelFunc
	err    error
}

func New(a *app.App) model {
	m := model{app: a}
	m.screen = m.newScreen(m.sessionScreen())
	m.ctx, m.cancel = context.WithCancel(context.Background())
	return m
}

func (m model) Init() tea.Cmd {
	return m.screen.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case message.ErrorMsg:
		m.err = msg.Err
		next, cmd, _ := m.screen.Update(msg)
		m.screen = next
		return m, cmd

	case message.SetServerURL:
		return m, m.setServerURL(msg.URL)

	case message.ServerURLSet:
		return m.navigate(m.sessionScreen())

	case message.Login:
		return m, m.login(msg.Username, msg.Password)

	case message.LoginResult:
		return m.navigate(m.sessionScreen())

	case message.SelectFolder:
		return m, m.selectFolder(msg.ID)

	case message.FolderSelected:
		return m.navigate(m.sessionScreen())

	case message.LoadFolders:
		return m, m.loadFolders()

	case message.FoldersMsg:
		next, cmd, _ := m.screen.Update(msg)
		m.screen = next
		return m, cmd

	case message.LoadAlbums:
		return m, m.loadAlbums()

	case message.AlbumsMsg:
		next, cmd, _ := m.screen.Update(msg)
		m.screen = next
		return m, cmd

	case message.TracksMsg:
		next, cmd, _ := m.screen.Update(msg)
		m.screen = next
		return m, cmd

	case screen.NavigateMsg:
		return m.navigate(msg.To)

	case tea.WindowSizeMsg:
		m.size = screen.FromWindowSize(msg)
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	next, cmd, err := m.screen.Update(msg)
	m.screen = next
	if err != nil {
		m.err = err
	}

	return m, cmd
}

func (m model) setServerURL(rawURL string) tea.Cmd {
	return func() tea.Msg {
		if err := m.app.SetServerURL(rawURL); err != nil {
			return message.ErrorMsg{Err: err}
		}
		return message.ServerURLSet{}
	}
}

func (m model) login(username, password string) tea.Cmd {
	return func() tea.Msg {
		if err := m.app.Login(m.ctx, username, password); err != nil {
			return message.ErrorMsg{Err: err}
		}
		return message.LoginResult{}
	}
}

func (m model) selectFolder(id string) tea.Cmd {
	return func() tea.Msg {
		if err := m.app.SelectFolder(id); err != nil {
			return message.ErrorMsg{Err: err}
		}
		return message.FolderSelected{}
	}
}

func (m model) loadFolders() tea.Cmd {
	return func() tea.Msg {
		folders, err := m.app.ListFolders(m.ctx)
		if err != nil {
			return message.ErrorMsg{Err: err}
		}
		return message.FoldersMsg{Folders: folders}
	}
}

func (m model) loadAlbums() tea.Cmd {
	return func() tea.Msg {
		albums, err := m.app.Albums(m.ctx)
		if err != nil {
			return message.ErrorMsg{Err: err}
		}
		return message.AlbumsMsg{Albums: albums}
	}
}

func (m model) loadTracks() tea.Cmd {
	return func() tea.Msg {
		tracks, err := m.app.Tracks(m.ctx, "")
		if err != nil {
			return message.ErrorMsg{Err: err}
		}
		return message.TracksMsg{Tracks: tracks}
	}
}

func (m model) sessionScreen() screen.Screen {
	switch m.app.RestoreSession() {
	case app.SessionReady:
		return screen.Library
	case app.SessionNeedsFolder:
		return screen.Folder
	case app.SessionNeedsLogin:
		return screen.Login
	default:
		return screen.Setup
	}
}

func (m model) View() tea.View {
	return m.screen.View(m.err, m.size)
}

func (m model) newScreen(to screen.Screen) screen.Model {
	switch to {
	case screen.Setup:
		return setup.New()
	case screen.Login:
		return login.New()
	case screen.Folder:
		return folder.New()
	case screen.Library:
		return library.New()
	default:
		return nil
	}
}

func (m model) navigate(to screen.Screen) (tea.Model, tea.Cmd) {
	if m.cancel != nil {
		m.cancel()
	}
	m.ctx, m.cancel = context.WithCancel(context.Background())

	m.screen = m.newScreen(to)
	m.err = nil

	return m, m.screen.Init()
}
