package urlinput

import tea "charm.land/bubbletea/v2"

func (m *UrlInput) Update(msg tea.Msg) tea.Cmd {
	methodCmd := m.method.Update(msg)
	textInputCmd := m.textInput.Update(msg)
	return tea.Batch(methodCmd, textInputCmd)
}
