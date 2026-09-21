package screens

import tea "charm.land/bubbletea/v2"

type Screen int

const (
	Setup Screen = iota
	Login
	Folder
	Library
)

type Size struct {
	Width  int
	Height int
}

func FromWindowSize(msg tea.WindowSizeMsg) Size {
	return Size{Width: msg.Width, Height: msg.Height}
}
