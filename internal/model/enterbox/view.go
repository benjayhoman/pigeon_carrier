package enterbox

import "charm.land/lipgloss/v2"

var (
	styleDim     = lipgloss.NewStyle().Faint(true)
	styleRegular = lipgloss.NewStyle()
)

func (b EnterBox) View() string {
	style := styleDim
	if b.focused {
		style = styleRegular
	}
	return style.Render("[" + b.symbol + "]")
}
