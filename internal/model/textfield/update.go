package textfield

import (
	tea "charm.land/bubbletea/v2"
)

func (m *TextField) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	var cmd tea.Cmd
	m.textArea, cmd = m.textArea.Update(msg)

	return cmd
}

func (m *TextField) Focus() tea.Cmd {
	m.focused = true
	return m.textArea.Focus()
}

func (m *TextField) Blur() tea.Cmd {
	m.focused = false
	m.textArea.Blur()
	return nil
}
