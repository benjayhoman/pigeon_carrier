package program

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	styleDim = lipgloss.NewStyle().Faint(true) // Dim for other protocols
)

func (m Program) View() tea.View {
	view := tea.NewView(fmt.Sprintf("%s\n%s\n%s%s",
		m.urlInput.View(),
		m.headers.View(),
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
