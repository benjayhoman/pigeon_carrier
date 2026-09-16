package headers

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
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
