package program

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const (
	logo = `
███████  ██  ██████  ███████  ██████  ███    ██ 
██   ██  ██ ██       ██      ██    ██ ████   ██ 
███████  ██ ██   ███ █████   ██    ██ ██ ██  ██ 
██       ██ ██    ██ ██      ██    ██ ██  ██ ██ 
██       ██  ██████  ███████  ██████  ██   ████

██████  █████   ██████  ██████  ██ ███████ ██████
██      ██   ██ ██   ██ ██   ██ ██ ██      ██   ██     
██ 🐦   ███████ ██████  ██████  ██ █████   ██████ 
██      ██   ██ ██   ██ ██   ██ ██ ██      ██   ██
██████  ██   ██ ██   ██ ██   ██ ██ ███████ ██   ██
`
)

var (
	styleDim = lipgloss.NewStyle().Faint(true) // Dim for other protocols
)

func (m Program) View() tea.View {
	view := tea.NewView(fmt.Sprintf("%s\n\n%s %s %s\n\n%s\n%s%s",
		logo,
		m.sendTab.View(),
		m.loadTab.View(),
		m.envTab.View(),
		m.request.View(),
		m.resolveResultsView(m.resultState),
		styleDim.Render("([Ctrl+s] to send, [Ctrl+q] to quit)")))
	view.AltScreen = true
	return view
}

func (m Program) resolveResultsView(resultsState ResultState) string {
	switch resultsState {
	case ResultStateHasResults:
		return m.results.View()
	case ResultStateIsSending:
		return m.sending.View()
	default:
		return ""
	}
}
