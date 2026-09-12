package component

import tea "charm.land/bubbletea/v2"

// Size represents the terminal dimensions available to a screen.
type Size struct {
	Width  int
	Height int
}

// FromWindowSize builds a Size from a tea.WindowSizeMsg.
func FromWindowSize(msg tea.WindowSizeMsg) Size {
	return Size{Width: msg.Width, Height: msg.Height}
}
