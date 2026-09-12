package setup

import (
	tea "charm.land/bubbletea/v2"

	"dolsh/ui/component"
	"dolsh/ui/login"
	"dolsh/ui/screen"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = component.FromWindowSize(msg)
		m.input = m.input.SetWidth(contentWidth(m.size.Width))
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			rawURL := m.input.Value()
			if err := m.app.SetServerURL(rawURL); err != nil {
				m.err = err
				return m, nil
			}

			return screen.Switch(m, login.New(m.app, m.size), tea.RequestWindowSize)
		default:
			m.err = nil
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}
