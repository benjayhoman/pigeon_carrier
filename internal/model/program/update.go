package program

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pigeon_carrier/internal/action"
)

func (m Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyTab:
			return m.incrementAndSetFocus(1)

		case tea.KeyShiftTab:
			return m.incrementAndSetFocus(-1)
		}

	case action.UpdateFocus:
		return m.incrementAndSetFocus(msg.Increment)
	}

	urlInputCmd := m.urlInput.Update(msg)
	methodCmd := m.method.Update(msg)
	headerCmd := m.headers.Update(msg)
	return m, tea.Batch(urlInputCmd, methodCmd, headerCmd)
}
