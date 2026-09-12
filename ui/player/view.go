package player

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"dolsh/ui/component"
)

func (m model) View() tea.View {
	style := lipgloss.NewStyle().
		Foreground(component.Muted).
		Bold(true)

	v := tea.NewView(style.Render("player — unimplemented"))
	v.AltScreen = true
	return v
}
