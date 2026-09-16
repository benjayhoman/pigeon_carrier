package program

import (
	"github.com/pigeon_carrier/internal/app"

	tea "github.com/charmbracelet/bubbletea"
)

// returns all focusable elements in the program so that the current focus can be managed
func (m *Program) focusables() []app.Focusable {
	focusables := make([]app.Focusable, 0, 3) // initial capacity for method, urlInput, and add header
	focusables = append(focusables, m.method, m.urlInput)

	for _, focusable := range m.headers.GetFocusables() {
		focusables = append(focusables, focusable)
	}
	return focusables
}

// calls the Focus method for what currently has focus and Blurs all others
func (m *Program) setFocus() (tea.Model, tea.Cmd) {
	return m.incrementAndSetFocus(0)
}

// increments the current focus by the given amount and updates the focus state of all focusable elements
func (m *Program) incrementAndSetFocus(increment int) (tea.Model, tea.Cmd) {
	focusables := m.focusables()
	length := len(focusables)

	if increment >= 0 {
		m.currentFocus = (m.currentFocus + increment) % length

	} else {
		m.currentFocus = (m.currentFocus + increment + length) % length
	}

	var cmd []tea.Cmd
	for i, f := range focusables {
		if i == m.currentFocus {
			cmd = append(cmd, f.Focus())

		} else {
			cmd = append(cmd, f.Blur())
		}
	}
	return m, tea.Batch(cmd...)
}
