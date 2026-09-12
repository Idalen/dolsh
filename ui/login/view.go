package login

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

	w := max(screenWidth-2*horizontalMargin, 20)

	return w
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

	labelStyle := lipgloss.NewStyle().
		Foreground(component.Foreground).
		Bold(true).
		Width(w).
		Align(lipgloss.Left)

	errorStyle := lipgloss.NewStyle().
		Foreground(component.Danger).
		Width(w).
		Align(lipgloss.Left)

	helpStyle := lipgloss.NewStyle().
		Foreground(component.Muted).
		Width(w).
		Align(lipgloss.Center)

	spacer := lipgloss.NewStyle().Width(w).Render("")

	labels := []string{"Username", "Password"}

	lines := []string{
		titleStyle.Render("dolsh"),
		subtitleStyle.Render("Sign in to " + m.app.ServerURL()),
		spacer,
	}

	for i, in := range m.input {
		lines = append(lines, labelStyle.Render(labels[i]))
		lines = append(lines, in.View())
		if i < len(m.input)-1 {
			lines = append(lines, spacer)
		}
	}

	lines = append(lines, spacer)

	if m.submitting {
		lines = append(lines, helpStyle.Render(m.spinner.View()+" signing in…"))
	} else {
		if m.err != nil {
			lines = append(lines, errorStyle.Render("✗ "+m.err.Error()))
		}
		lines = append(lines, spacer, helpStyle.Render("enter · sign in      ↑/↓ · switch      q · quit"))
	}

	block := lipgloss.JoinVertical(lipgloss.Left, lines...)

	view := block
	if m.size.Width > 0 && m.size.Height > 0 {
		view = lipgloss.Place(m.size.Width, m.size.Height, lipgloss.Center, lipgloss.Center, block)
	}

	v := tea.NewView(view)
	v.AltScreen = true

	return v
}
