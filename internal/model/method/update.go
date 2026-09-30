package method

import tea "charm.land/bubbletea/v2"

func (m *Method) Update(msg tea.Msg) tea.Cmd {
	return m.option.Update(msg)
}

func (m *Method) Focus() tea.Cmd {
	return m.option.Focus()
}

func (m *Method) Blur() tea.Cmd {
	return m.option.Blur()
}
