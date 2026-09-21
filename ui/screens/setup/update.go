package setup

import (
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
		m.err = msg.Err
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "enter":
			return m, m.submit(m.input.Value())
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m model) submit(rawURL string) tea.Cmd {
	return func() tea.Msg {
		if err := m.svc.SetServerURL(rawURL); err != nil {
			return message.Error{Err: err}
		}
		return message.Next{}
	}
}
