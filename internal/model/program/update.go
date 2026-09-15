package program

import tea "github.com/charmbracelet/bubbletea"

type focusable interface {
	Focus() tea.Cmd
	Blur() tea.Cmd
}

func (m Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyTab:
			m.currentFocus = (m.currentFocus + 1) % len(m.focusables())
			return m.updateFocus()

		case tea.KeyShiftTab:
			m.currentFocus = (m.currentFocus - 1 + len(m.focusables())) % len(m.focusables())
			return m.updateFocus()
		}
	}

	urlInputCmd := m.urlInput.Update(msg)
	methodCmd := m.method.Update(msg)
	return m, tea.Batch(urlInputCmd, methodCmd)
}

func (m *Program) updateFocus() (tea.Model, tea.Cmd) {
	var cmd []tea.Cmd

	for i, f := range m.focusables() {
		if i == m.currentFocus {
			cmd = append(cmd, f.Focus())

		} else {
			cmd = append(cmd, f.Blur())
		}
	}

	return m, tea.Batch(cmd...)
}

func (m *Program) focusables() []focusable {
	return []focusable{m.urlInput, m.method}
}
