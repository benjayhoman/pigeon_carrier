package urlinput

import (
	"fmt"
	"regexp"
	"strings"

	"charm.land/lipgloss/v2"
)

var (
	styleHTTPS = lipgloss.NewStyle().Foreground(lipgloss.Color("10")) // Green for https
	styleOther = lipgloss.NewStyle().Faint(true)                      // Dim for other protocols
	styleHost  = lipgloss.NewStyle().Foreground(lipgloss.Color("13")) // Purple for hostname
	styleQuery = lipgloss.NewStyle().Foreground(lipgloss.Color("14")) // Cyan for query params
	stylePath  = lipgloss.NewStyle()                                  // Normal for path
	protoRegex = regexp.MustCompile(`^(?i)([a-z][a-z0-9+.-]*:(?://)?)`)
)

func (m UrlInput) View() string {
	var inputView string
	val := m.textInput.Value()

	if strings.TrimSpace(val) != "" {
		highlighted := m.highlightWithCursor(val, m.textInput.Position(), styleCursor)
		inputView = m.textInput.Prompt + highlighted

	} else {
		inputView = m.textInput.View()
	}

	return fmt.Sprintf(
		"%s %s",
		m.method.View(),
		inputView,
	)
}

// TODO: clean this up for readability
func (m *UrlInput) highlightWithCursor(raw string, pos int, cursorStyle lipgloss.Style) string {
	if raw == "" {
		return cursorStyle.Render(" ")
	}

	runes := []rune(raw)
	if pos < 0 {
		pos = 0
	}
	if pos > len(runes) {
		pos = len(runes)
	}

	// Token types
	const (
		tokDefault = iota
		tokHTTPS
		tokOtherProto
		tokHost
		tokQuery
	)

	tokenMap := make([]int, len(runes))

	// 1. Identify protocol
	if match := protoRegex.FindString(raw); match != "" {
		matchRunes := []rune(match)
		pType := tokOtherProto
		if strings.HasPrefix(strings.ToLower(match), "https") {
			pType = tokHTTPS
		}
		for i := 0; i < len(matchRunes) && i < len(tokenMap); i++ {
			tokenMap[i] = pType
		}
	}

	// 2. Identify query & fragment
	queryStart := strings.IndexAny(raw, "?#")
	if queryStart != -1 {
		qStartRune := len([]rune(raw[:queryStart]))
		for i := qStartRune; i < len(tokenMap); i++ {
			tokenMap[i] = tokQuery
		}
	}

	// 3. Identify host
	protoMatch := protoRegex.FindString(raw)
	protoLenRunes := len([]rune(protoMatch))

	preQuery := raw
	if queryStart != -1 {
		preQuery = raw[:queryStart]
	}

	if len(preQuery) > len(protoMatch) {
		afterProto := preQuery[len(protoMatch):]
		pathIdx := strings.IndexByte(afterProto, '/')
		hostEndByte := len(preQuery)
		if pathIdx != -1 {
			hostEndByte = len(protoMatch) + pathIdx
		}
		hostEndRune := len([]rune(raw[:hostEndByte]))
		for i := protoLenRunes; i < hostEndRune && i < len(tokenMap); i++ {
			tokenMap[i] = tokHost
		}
	}

	// Render each token group, handling the cursor position seamlessly
	var b strings.Builder
	for i := 0; i <= len(runes); i++ {
		if i == pos && m.focused {
			cursorChar := " "
			if i < len(runes) {
				cursorChar = string(runes[i])
			}
			b.WriteString(cursorStyle.Render(cursorChar))
			if i == len(runes) {
				break
			}
			continue
		}

		if i >= len(runes) {
			break
		}

		ch := string(runes[i])
		switch tokenMap[i] {
		case tokHTTPS:
			b.WriteString(styleHTTPS.Render(ch))
		case tokOtherProto:
			b.WriteString(styleOther.Render(ch))
		case tokHost:
			b.WriteString(styleHost.Render(ch))
		case tokQuery:
			b.WriteString(styleQuery.Render(ch))
		default:
			b.WriteString(stylePath.Render(ch))
		}
	}

	return b.String()
}
