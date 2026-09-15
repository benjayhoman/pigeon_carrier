package urlinput

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m *UrlInput) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return cmd
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
