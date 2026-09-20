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

// NextMsg asks the root model to advance to the screen determined by the
// current session state.
type NextMsg struct{}

func Next() tea.Cmd {
	return func() tea.Msg {
		return NextMsg{}
	}
}
