package option

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

var styleDim = lipgloss.NewStyle().Faint(true)

func (o Option) View(value string) string {
	if o.focused {
		return fmt.Sprintf("[↑↓ %s]", value)
	}
	return fmt.Sprintf("%s%s%s", styleDim.Render("["), value, styleDim.Render("]"))
}
