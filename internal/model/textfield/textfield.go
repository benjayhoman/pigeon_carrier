package textfield

import (
	"charm.land/bubbles/v2/textarea"
	tea "charm.land/bubbletea/v2"
)

type TextField struct {
	textArea textarea.Model
	focused  bool
}

func NewTextField() *TextField {
	input := textarea.New()
	input.Prompt = ""
	input.ShowLineNumbers = true
	input.SetVirtualCursor(true)

	return &TextField{textArea: input}
}

func (m TextField) Init() tea.Cmd {
	return nil
}

func (m TextField) Value() string {
	return m.textArea.Value()
}

func (m *TextField) SetValue(value string) {
	m.textArea.SetValue(value)
}
