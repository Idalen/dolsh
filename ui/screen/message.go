package screen

import tea "charm.land/bubbletea/v2"

type Screen int

const (
	Setup Screen = iota
	Login
	Folder
	Library
)

type NavigateMsg struct {
	To Screen
}

func Go(to Screen) tea.Cmd {
	return func() tea.Msg {
		return NavigateMsg{To: to}
	}
}
