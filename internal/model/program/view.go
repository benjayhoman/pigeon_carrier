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
	urlInputView := m.urlInput.View()
	headersView := m.headers.View()
	resultsView := m.results.View()

	view := tea.NewView(fmt.Sprintf("%s\n%s\n%s%s",
		urlInputView,
		headersView,
		resultsView,
		styleDim.Render("([Ctrl+s] to send, [Ctrl+q] to quit)")))
	view.AltScreen = true
	return view
}
