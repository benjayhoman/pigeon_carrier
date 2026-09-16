package addremovebox

import "charm.land/lipgloss/v2"

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
