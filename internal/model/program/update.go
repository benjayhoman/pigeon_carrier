package program

import (
	"github.com/pigeon_carrier/internal/action"

	tea "charm.land/bubbletea/v2"
)

func (m Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab":
			return m.incrementAndSetFocus(1)

		case "shift+tab":
			return m.incrementAndSetFocus(-1)
		}

	case action.UpdateFocus:
		return m.incrementAndSetFocus(msg.Increment)
	}

	urlInputCmd := m.urlInput.Update(msg)
	methodCmd := m.method.Update(msg)
	headerCmd := m.headers.Update(msg)
	resultsCmd := m.results.Update(msg)
	return m, tea.Batch(urlInputCmd, methodCmd, headerCmd, resultsCmd)
}
