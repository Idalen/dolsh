package component

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const inputBoxChrome = 4

type Input struct {
	textinput.Model
	width int
}

func (i Input) Update(msg tea.Msg) (Input, tea.Cmd) {
	var cmd tea.Cmd
	i.Model, cmd = i.Model.Update(msg)
	return i, cmd
}

func (i Input) SetWidth(width int) Input {
	i.width = width
	i.Model.SetWidth(max(width-inputBoxChrome, 1))
	return i
}

func (i Input) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1)

	if i.Focused() {
		style = style.BorderForeground(Accent)
	} else {
		style = style.BorderForeground(BorderIdle)
	}

	return style.Render(i.Model.View())
}

func NewInput(placeholder string) Input {
	return newInput(placeholder, false)
}

func NewPasswordInput(placeholder string) Input {
	return newInput(placeholder, true)
}

func newInput(placeholder string, password bool) Input {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Prompt = ""
	ti.CharLimit = 0
	ti.SetStyles(inputStyles())
	ti.Focus()

	if password {
		ti.EchoMode = textinput.EchoPassword
		ti.EchoCharacter = '•'
	}

	in := Input{Model: ti}
	return in.SetWidth(60)
}

func inputStyles() textinput.Styles {
	styles := textinput.DefaultStyles(true)

	styles.Focused.Text = lipgloss.NewStyle().Foreground(Foreground)
	styles.Focused.Placeholder = lipgloss.NewStyle().Foreground(Muted)
	styles.Focused.Prompt = lipgloss.NewStyle().Foreground(Accent)

	styles.Blurred.Text = lipgloss.NewStyle().Foreground(Foreground)
	styles.Blurred.Placeholder = lipgloss.NewStyle().Foreground(Muted)
	styles.Blurred.Prompt = lipgloss.NewStyle().Foreground(Muted)

	styles.Cursor = textinput.CursorStyle{
		Color: Accent,
		Shape: tea.CursorBar,
		Blink: true,
	}

	return styles
}
