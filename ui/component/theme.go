// Package component provides reusable UI building blocks for dolsh screens.
package component

import "charm.land/lipgloss/v2"

// Shared color palette for the dolsh UI. Jellyfin-inspired violet accent on a
// dark, low-contrast background.
var (
	// Accent is the primary brand color.
	Accent = lipgloss.Color("#A78BFA")
	// AccentSoft is a lighter variant of the accent.
	AccentSoft = lipgloss.Color("#C4B5FD")
	// Muted is used for de-emphasized text.
	Muted = lipgloss.Color("#8B8B9E")
	// Foreground is the default text color.
	Foreground = lipgloss.Color("#EDEDF7")
	// Danger marks error states.
	Danger = lipgloss.Color("#FF6B6B")
	// BorderActive is the border color for the focused element.
	BorderActive = lipgloss.Color("#A78BFA")
	// BorderIdle is the border color for unfocused elements.
	BorderIdle = lipgloss.Color("#3A3A4A")
)
