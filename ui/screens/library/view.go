package library

import (
	"fmt"
	"strings"
	"time"

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

func formatDuration(d time.Duration) string {
	total := int(d.Round(time.Second).Seconds())
	h := total / 3600
	m := (total % 3600) / 60
	s := total % 60
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%d:%02d", m, s)
}

func (m model) View() tea.View {
	w := contentWidth(m.size.Width)
	gap := 2
	colW := (w - gap) / 2

	titleStyle := lipgloss.NewStyle().
		Foreground(component.Accent).
		Bold(true).
		Width(w).
		Align(lipgloss.Center)

	colTitleStyle := lipgloss.NewStyle().
		Foreground(component.Accent).
		Bold(true).
		Width(colW).
		Align(lipgloss.Left)

	itemStyle := lipgloss.NewStyle().
		Foreground(component.Foreground)

	selectedStyle := lipgloss.NewStyle().
		Foreground(component.Accent).
		Bold(true)

	mutedStyle := lipgloss.NewStyle().
		Foreground(component.Muted).
		Faint(true)

	errorStyle := lipgloss.NewStyle().
		Foreground(component.Danger).
		Width(w).
		Align(lipgloss.Left)

	helpStyle := lipgloss.NewStyle().
		Foreground(component.Muted).
		Width(w).
		Align(lipgloss.Center)

	spacer := lipgloss.NewStyle().Width(w).Render("")

	left := []string{colTitleStyle.Render("Albums")}
	for i, album := range m.albums {
		prefix := "  "
		nameStyle := itemStyle
		if i == m.selector {
			prefix = "▸ "
			nameStyle = selectedStyle
		}
		name := prefix + album.Name
		pad := colW - len([]rune(name)) - len([]rune(album.Artist))
		if pad < 1 {
			pad = 1
		}
		left = append(left, nameStyle.Render(name)+strings.Repeat(" ", pad)+mutedStyle.Render(album.Artist))
	}

	rightTitle := "Tracks"
	var selectedAlbumID string
	if album, ok := m.selectedAlbum(); ok {
		selectedAlbumID = album.ID
		rightTitle = album.Name
		if album.Year > 0 {
			rightTitle = fmt.Sprintf("%s (%d)", album.Name, album.Year)
		}
	}
	right := []string{colTitleStyle.Render(rightTitle)}
	for _, track := range m.tracks[selectedAlbumID] {
		name := fmt.Sprintf("%2d. %s", track.Index, track.Name)
		dur := formatDuration(track.Duration)
		pad := colW - len([]rune(name)) - len([]rune(dur))
		if pad < 1 {
			pad = 1
		}
		right = append(right, itemStyle.Render(name)+strings.Repeat(" ", pad)+mutedStyle.Render(dur))
	}

	columns := lipgloss.JoinHorizontal(
		lipgloss.Top,
		lipgloss.JoinVertical(lipgloss.Left, left...),
		lipgloss.NewStyle().Width(gap).Render(""),
		lipgloss.JoinVertical(lipgloss.Left, right...),
	)

	lines := []string{
		titleStyle.Render("dolsh"),
		spacer,
		columns,
	}

	if m.err != nil {
		lines = append(lines, errorStyle.Render("✗ "+m.err.Error()))
	}

	lines = append(lines, spacer, helpStyle.Render("↑/↓ · choose      q · quit"))

	block := lipgloss.JoinVertical(lipgloss.Left, lines...)

	view := block
	if m.size.Width > 0 && m.size.Height > 0 {
		view = lipgloss.Place(m.size.Width, m.size.Height, lipgloss.Center, lipgloss.Center, block)
	}

	v := tea.NewView(view)
	v.AltScreen = true

	return v
}
