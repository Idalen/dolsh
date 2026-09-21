package ui

import (
	"context"

	tea "charm.land/bubbletea/v2"

	"dolsh/app"
	"dolsh/ui/message"
	"dolsh/ui/screens"
	"dolsh/ui/screens/folder"
	"dolsh/ui/screens/library"
	"dolsh/ui/screens/login"
	"dolsh/ui/screens/setup"
)

type model struct {
	screen tea.Model

	app    *app.App
	size   screens.Size
	ctx    context.Context
	cancel context.CancelFunc
}

func New(a *app.App) model {
	m := model{app: a}
	m.ctx, m.cancel = context.WithCancel(context.Background())
	m.screen = m.newScreen(m.sessionScreen())
	return m
}

func (m model) Init() tea.Cmd {
	return m.screen.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case message.Next:
		return m.navigate(m.sessionScreen())

	case tea.WindowSizeMsg:
		m.size = screens.FromWindowSize(msg)
		next, cmd := m.screen.Update(msg)
		m.screen = next
		return m, cmd

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	next, cmd := m.screen.Update(msg)
	m.screen = next

	return m, cmd
}

func (m model) sessionScreen() screens.Screen {
	switch m.app.RestoreSession() {
	case app.SessionReady:
		return screens.Library
	case app.SessionNeedsFolder:
		return screens.Folder
	case app.SessionNeedsLogin:
		return screens.Login
	default:
		return screens.Setup
	}
}

func (m model) View() tea.View {
	return m.screen.View()
}

func (m model) newScreen(to screens.Screen) tea.Model {
	switch to {
	case screens.Setup:
		return setup.New(m.app)
	case screens.Login:
		return login.New(m.app, m.ctx)
	case screens.Folder:
		return folder.New(m.app, m.ctx)
	case screens.Library:
		return library.New(m.app, m.ctx)
	default:
		return nil
	}
}

func (m model) navigate(to screens.Screen) (tea.Model, tea.Cmd) {
	if m.cancel != nil {
		m.cancel()
	}
	m.ctx, m.cancel = context.WithCancel(context.Background())

	m.screen = m.newScreen(to)

	return m, tea.Batch(
		m.screen.Init(),
		func() tea.Msg {
			return tea.WindowSizeMsg{Width: m.size.Width, Height: m.size.Height}
		},
	)
}
