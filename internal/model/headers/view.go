package headers

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

var (
	styleDim = lipgloss.NewStyle().Faint(true)
)

func (h Headers) View() string {
	result := fmt.Sprintf("%s %s\n", styleDim.Render("Headers"), h.addNewHeader.View())
	for _, header := range h.headers {
		result += fmt.Sprintf("  %s\n", header.View())
	}
	return result
}
