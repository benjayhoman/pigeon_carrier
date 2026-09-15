package method

import tea "github.com/charmbracelet/bubbletea"

func (m *Method) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			m.selected = (m.selected - 1 + len(listTypes())) % len(listTypes())

		case tea.KeyDown:
			m.selected = (m.selected + 1) % len(listTypes())
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
