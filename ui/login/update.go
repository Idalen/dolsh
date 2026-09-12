package login

import (
	"log"

	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/component"
	"dolsh/ui/folder"
	"dolsh/ui/screen"
)

type loginMsg struct {
	err error
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loginMsg:
		if msg.err != nil {
			log.Printf("login failed: %v", msg.err)
			m.err = msg.err
			m.submitting = false
			return m, nil
		}
		log.Println("login succeeded")
		return screen.Switch(m, library.New(m.app, m.size))

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.WindowSizeMsg:
		return m.withSize(component.FromWindowSize(msg)), nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.Cancel()
			return m, tea.Quit
		}

		if m.submitting {
			return m, nil
		}

		switch msg.String() {
		case "up":
			m.input[m.selector].Blur()
			m.selector = (m.selector - 1 + len(m.input)) % len(m.input)
			return m, m.input[m.selector].Focus()
		case "down":
			m.input[m.selector].Blur()
			m.selector = (m.selector + 1) % len(m.input)
			return m, m.input[m.selector].Focus()
		case "enter":
			username := m.input[0].Value()
			password := m.input[1].Value()
			m.submitting = true
			m.err = nil
			return m, tea.Batch(m.login(username, password), m.spinTick())
		}
	}

	if m.submitting {
		return m, nil
	}

	var cmd tea.Cmd
	m.input[m.selector], cmd = m.input[m.selector].Update(msg)

	return m, cmd
}

func (m model) spinTick() tea.Cmd {
	return func() tea.Msg { return m.spinner.Tick() }
}

func (m model) login(username, password string) tea.Cmd {
	return func() tea.Msg {
		log.Printf("login attempt for user %q", username)
		err := m.app.Login(m.ctx, username, password)
		return loginMsg{err: err}
	}
}
