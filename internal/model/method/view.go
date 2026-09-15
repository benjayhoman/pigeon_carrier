package method

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleGET    = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))  // Green for get
	stylePOST   = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))  // Yellow for post
	stylePUT    = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))  // Cyan for put
	stylePATCH  = lipgloss.NewStyle().Foreground(lipgloss.Color("13"))  // Magenta for patch
	styleDELETE = lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red for delete
	styleDim    = lipgloss.NewStyle().Faint(true)
)

func (m Method) View() string {
	method := listTypes()[m.selected]
	style := getStyleForMethod(method)

	if m.focused {
		return fmt.Sprintf("[↑↓ %s]", style.Render(string(method)))

	} else {
		return fmt.Sprintf("%s%s%s", styleDim.Render("["), style.Render(string(method)), styleDim.Render("]"))
	}
}

func getStyleForMethod(method MethodType) lipgloss.Style {
	switch method {
	case GET:
		return styleGET
	case POST:
		return stylePOST
	case PUT:
		return stylePUT
	case PATCH:
		return stylePATCH
	case DELETE:
		return styleDELETE
	default:
		return styleDim
	}
}
