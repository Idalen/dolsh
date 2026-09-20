package login

import (
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/message"
	"dolsh/ui/screen"
)

func (m model) Update(msg tea.Msg) (screen.Model, tea.Cmd, error) {
	switch msg := msg.(type) {
	case message.ErrorMsg:
		m.submitting = false
		return m, nil, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd, nil

	case tea.KeyPressMsg:
		if m.submitting {
			return m, nil, nil
		}

		switch msg.String() {
		case "up":
			m.input[m.selector].Blur()
			m.selector = (m.selector - 1 + len(m.input)) % len(m.input)
			return m, m.input[m.selector].Focus(), nil
		case "down":
			m.input[m.selector].Blur()
			m.selector = (m.selector + 1) % len(m.input)
			return m, m.input[m.selector].Focus(), nil
		case "enter":
			username := m.input[0].Value()
			password := m.input[1].Value()
			m.submitting = true
			return m, tea.Batch(message.SubmitLogin(username, password), m.spinTick()), nil
		}
	}

	var cmd tea.Cmd
	m.input[m.selector], cmd = m.input[m.selector].Update(msg)

	return m, cmd, nil
}

func (m model) spinTick() tea.Cmd {
	return func() tea.Msg { return m.spinner.Tick() }
}
