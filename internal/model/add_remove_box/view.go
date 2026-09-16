package addremovebox

import "github.com/charmbracelet/lipgloss"

var (
	styleDim     = lipgloss.NewStyle().Faint(true)
	styleRegular = lipgloss.NewStyle()
)

func (b AddRemoveBox) View() string {
	style := styleDim
	if b.focused {
		style = styleRegular
	}
	return style.Render("[" + b.symbol + "]")
}
