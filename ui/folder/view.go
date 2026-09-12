package library

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"dolsh/ui/component"
)

const (
	defaultWidth     = 40
	horizontalMargin = 20
)

func contentWidth(screenWidth int) int {
	if screenWidth <= 0 {
		return defaultWidth
	}

	return max(screenWidth-2*horizontalMargin, 20)
}

func (m model) View() tea.View {
	w := contentWidth(m.size.Width)

	titleStyle := lipgloss.NewStyle().
		Foreground(component.Accent).
		Bold(true).
		Width(w).
		Align(lipgloss.Center)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(component.Muted).
		Width(w).
		Align(lipgloss.Center)

	errorStyle := lipgloss.NewStyle().
		Foreground(component.Danger).
		Width(w).
		Align(lipgloss.Left)

	helpStyle := lipgloss.NewStyle().
		Foreground(component.Muted).
		Width(w).
		Align(lipgloss.Center)

	spacer := lipgloss.NewStyle().Width(w).Render("")

	lines := []string{
		titleStyle.Render("dolsh"),
		subtitleStyle.Render("Choose a library"),
		spacer,
	}

	if m.err != nil {
		lines = append(lines, errorStyle.Render("✗ "+m.err.Error()))
	}

	lines = append(lines, m.renderFolders(w))

	lines = append(lines, spacer, helpStyle.Render("↑/↓ · choose      enter · open      q · quit"))

	block := lipgloss.JoinVertical(lipgloss.Left, lines...)

	view := block
	if m.size.Width > 0 && m.size.Height > 0 {
		view = lipgloss.Place(m.size.Width, m.size.Height, lipgloss.Center, lipgloss.Center, block)
	}

	v := tea.NewView(view)
	v.AltScreen = true

	return v
}

func (m model) renderFolders(w int) string {
	itemStyle := lipgloss.NewStyle().
		Foreground(component.Foreground).
		Width(w).
		Align(lipgloss.Left)

	selectedStyle := lipgloss.NewStyle().
		Foreground(component.Accent).
		Bold(true).
		Width(w).
		Align(lipgloss.Left)

	lines := make([]string, 0, len(m.folders))
	for i, folder := range m.folders {
		prefix := "  "
		style := itemStyle
		if i == m.selector {
			prefix = "▸ "
			style = selectedStyle
		}
		lines = append(lines, style.Render(prefix+folder.Name))
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
