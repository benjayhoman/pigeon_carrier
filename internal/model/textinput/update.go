package textinput

import (
	"github.com/pigeon_carrier/internal/model/common"

	tea "charm.land/bubbletea/v2"
)

func (m *TextInput) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	cmd := common.HandleCopyAndPaste(msg, &m.input)
	var updateCmd tea.Cmd
	m.input, updateCmd = m.input.Update(msg)
	return tea.Batch(cmd, updateCmd)
}

func (m *TextInput) Focus() tea.Cmd {
	m.focused = true
	return m.input.Focus()
}

func (m *TextInput) Blur() tea.Cmd {
	m.focused = false
	m.input.Blur()
	return nil
}
