package textinput

import (
	bubblesinput "charm.land/bubbles/v2/textinput"
	"charm.land/lipgloss/v2"
	"github.com/pigeon_carrier/internal/app"
)

var stylePrompt = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))

type TextInput struct {
	input   bubblesinput.Model
	focused bool
}

func NewTextInput(placeholder string, charLimit int, width int) *TextInput {
	input := bubblesinput.New()
	input.Placeholder = placeholder
	input.SetVirtualCursor(true)
	input.Focus()
	input.CharLimit = charLimit
	input.SetWidth(width)
	styles := input.Styles()
	styles.Focused.Prompt = stylePrompt
	styles.Blurred.Prompt = stylePrompt
	styles.Cursor.Color = lipgloss.Color("15")
	input.SetStyles(styles)

	return &TextInput{input: input}
}

func (m *TextInput) Value() string {
	return m.input.Value()
}

func (m *TextInput) Position() int {
	return m.input.Position()
}

func (m *TextInput) Prompt() string {
	return m.input.Prompt
}

func (m *TextInput) IsFocused() bool {
	return m.focused
}

func (m *TextInput) GetFocusables() []app.Focusable {
	return []app.Focusable{m}
}
