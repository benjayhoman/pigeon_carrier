package textfield

import (
	"encoding/json"

	tea "charm.land/bubbletea/v2"
	"github.com/tidwall/pretty"
)

func (m *TextField) Update(msg tea.Msg) tea.Cmd {
	if !m.focused {
		return nil
	}

	var cmd tea.Cmd
	m.textArea, cmd = m.textArea.Update(msg)

	jsonFormatted, err := formatBodyAsJson(m.textArea.Value())
	if err == nil {
		m.textArea.SetValue(jsonFormatted)
	}

	return cmd
}

func (m *TextField) Focus() tea.Cmd {
	m.focused = true
	return m.textArea.Focus()
}

func (m *TextField) Blur() tea.Cmd {
	m.focused = false
	m.textArea.Blur()
	return nil
}

func formatBodyAsJson(body string) (string, error) {
	var obj interface{}
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		return "Error parsing JSON", err
	}

	opts := pretty.Options{
		Width:  15,
		Indent: "  ",
	}

	formattedJSON := pretty.PrettyOptions([]byte(body), &opts)
	colorizedJSON := pretty.Color(formattedJSON, nil)

	// 3. Convert to string and print
	result := string(colorizedJSON)
	return result, nil
}
