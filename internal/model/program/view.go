package program

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleDim = lipgloss.NewStyle().Faint(true) // Dim for other protocols
)

func (m Program) View() string {
	urlInputView := m.urlInput.View()
	headersView := m.headers.View()
	return fmt.Sprintf("%s\n%s\n%s",
		urlInputView,
		headersView,
		styleDim.Render("([Ctrl+Enter] to send, [Ctrl+C] to quit)"))
}
