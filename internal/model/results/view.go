package results

import (
	"fmt"
	"sort"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	styleDim = lipgloss.NewStyle().Faint(true)                     // Dim for other protocols
	green    = lipgloss.NewStyle().Foreground(lipgloss.Color("2")) // Green for success
	blue     = lipgloss.NewStyle().Foreground(lipgloss.Color("4")) // Blue for informational
	red      = lipgloss.NewStyle().Foreground(lipgloss.Color("1")) // Red for errors
	wrapText = lipgloss.NewStyle().Width(80)
)

func (r Results) View() string {
	switch r.state {
	case ResultStateNone:
		return ""

	case ResultStateIsSending:
		return fmt.Sprintf("%s %s \n\n", "", sendAnimation[r.sendAnimationFrame])

	case ResultStateHasResults:
		return fmt.Sprintf("%s %s %s\n%s\n%s\n\n",
			styleDim.Render("Response:"),
			styleStatus(fmt.Sprintf("%d", r.Status)),
			r.toggleHeadersButton.View(),
			r.styleHeaders(r.Headers),
			wrapText.Render(r.Body))
	}
	return ""
}

func (r Results) styleHeaders(headers map[string]string) string {
	if len(headers) == 0 || r.hideHeaders {
		return ""
	}

	keys := make([]string, 0, len(headers))
	for k := range headers {
		keys = append(keys, k)
	}

	// 2. Sort the keys alphabetically
	sort.Strings(keys)

	var sb strings.Builder
	for _, key := range keys {
		value := headers[key]
		sb.WriteString(fmt.Sprintf("%s : %s\n", styleDim.Render(key), styleDim.Render(value)))
	}
	return sb.String()
}

func styleStatus(status string) string {
	if strings.HasPrefix(status, "2") {
		return green.Render(status)

	} else if strings.HasPrefix(status, "4") {
		return blue.Render(status)

	} else if strings.HasPrefix(status, "5") {
		return red.Render(status)

	} else {
		return status
	}
}
