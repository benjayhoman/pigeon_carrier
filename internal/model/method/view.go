package method

import (
	"charm.land/lipgloss/v2"
)

var (
	styleGET    = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	stylePOST   = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
	stylePUT    = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	stylePATCH  = lipgloss.NewStyle().Foreground(lipgloss.Color("13"))
	styleDELETE = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	styleDim    = lipgloss.NewStyle().Faint(true)
)

func (m Method) View() string {
	method := MethodType(m.option.Value())
	style := getStyleForMethod(method)
	return m.option.View(style.Render(string(method)))
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
