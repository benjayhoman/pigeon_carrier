package results

import (
	"encoding/json"

	"github.com/pigeon_carrier/internal/action"

	tea "charm.land/bubbletea/v2"
	"github.com/tidwall/pretty"
)

func (r *Results) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.body.SetWidth(msg.Width)
		return nil

	case action.UpdateResults:
		jsonBody, err := formatBodyAsJson(msg.Body)
		if err != nil {
			r.body.SetContent(msg.Body)
			r.body.SetHeight(20)

		} else {
			r.body.SetContent(jsonBody)
			r.body.SetHeight(20)
		}
		r.Status = msg.StatusCode
		r.Headers = msg.Headers
		r.hideHeaders = true
		return nil

	case action.ToggleResultHeaders:
		r.hideHeaders = !r.hideHeaders
		if r.hideHeaders {
			r.toggleHeadersButton.SetSymbol("Show Headers")
		} else {
			r.toggleHeadersButton.SetSymbol("Hide Headers")
		}
		return nil
	}

	var cmd tea.Cmd
	r.body, cmd = r.body.Update(msg)
	return tea.Batch(cmd, r.toggleHeadersButton.Update(msg))
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
