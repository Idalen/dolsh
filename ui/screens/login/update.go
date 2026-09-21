package login

import (
	"charm.land/bubbles/v2/spinner"
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
		m.submitting = false
		m.err = msg.Err
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyPressMsg:
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
			return m, tea.Batch(m.submit(username, password), m.spinTick())
		}
	}

	var cmd tea.Cmd
	m.input[m.selector], cmd = m.input[m.selector].Update(msg)

	return m, cmd
}

func (m model) spinTick() tea.Cmd {
	return func() tea.Msg { return m.spinner.Tick() }
}

func (m model) submit(username, password string) tea.Cmd {
	return func() tea.Msg {
		if err := m.svc.Login(m.ctx, username, password); err != nil {
			return message.Error{Err: err}
		}
		return message.Next{}
	}
}
