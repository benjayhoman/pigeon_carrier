package urlinput

import (
	"github.com/pigeon_carrier/internal/model/common"

	tea "charm.land/bubbletea/v2"
)

func (m *UrlInput) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	cmd := common.HandleCopyAndPaste(msg, &m.textInput)

	var updateCmd tea.Cmd
	m.textInput, updateCmd = m.textInput.Update(msg)
	return tea.Batch(cmd, updateCmd)
}

func (m *UrlInput) Focus() tea.Cmd {
	m.focused = true
	return m.textInput.Focus()
}

func (m *UrlInput) Blur() tea.Cmd {
	m.focused = false
	m.textInput.Blur()
	return nil
}
