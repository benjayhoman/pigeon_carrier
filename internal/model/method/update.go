package method

import tea "charm.land/bubbletea/v2"

func (m *Method) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up":
			m.selected = (m.selected - 1 + listedTypesLen) % listedTypesLen

		case "down":
			m.selected = (m.selected + 1) % listedTypesLen
		}
	}
	return nil
}

func (m *Method) Focus() tea.Cmd {
	m.focused = true
	return nil
}

func (m *Method) Blur() tea.Cmd {
	m.focused = false
	return nil
}
