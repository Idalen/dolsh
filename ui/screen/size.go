package screen

import tea "charm.land/bubbletea/v2"

type Size struct {
	Width  int
	Height int
}

func FromWindowSize(msg tea.WindowSizeMsg) Size {
	return Size{Width: msg.Width, Height: msg.Height}
}
