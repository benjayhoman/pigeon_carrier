package scrollabletext

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	styleDim    = lipgloss.NewStyle().Faint(true) // Dim for other protocols
	styleNormal = lipgloss.NewStyle()
)

func (r ScrollableText) View() string {
	if !r.focused {
		return styleDim.Render(r.viewport.View())
	}
	return styleNormal.Render(r.viewport.View())
}
