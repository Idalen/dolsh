package screen

import (
	tea "charm.land/bubbletea/v2"
)

type Model interface {
	Init() tea.Cmd
	Update(tea.Msg) (Model, tea.Cmd, error)
	View(error, Size) tea.View
}

type Cancellable interface {
	Cancel()
}

func Switch(current, next tea.Model, extra ...tea.Cmd) (tea.Model, tea.Cmd) {
	if c, ok := current.(Cancellable); ok {
		c.Cancel()
	}

	cmds := append([]tea.Cmd{next.Init()}, extra...)
	return next, tea.Batch(cmds...)
}
