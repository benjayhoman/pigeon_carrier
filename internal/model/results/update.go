package results

import (
	"encoding/json"
	"time"

	"github.com/pigeon_carrier/internal/action"

	tea "charm.land/bubbletea/v2"
	"github.com/tidwall/pretty"
)

var (
	sendAnimation = []string{
		"====>",
		" ====>",
		"  ====>",
		"   ====>",
		"    ====>",
		"     ====>",
		"      ====>",
		"       ====>",
		"        ====>",
		"        <====",
		"       <====",
		"      <====",
		"     <====",
		"    <====",
		"   <====",
		"  <====",
		" <====",
		"<====",
	}
)

type tickMsg time.Time

func (r *Results) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case action.UpdateResults:
		jsonBody, err := formatBodyAsJson(msg.Body)
		if err != nil {
			r.Body = msg.Body

		} else {
			r.Body = jsonBody
		}
		r.Status = msg.StatusCode
		r.Headers = msg.Headers
		r.hideHeaders = false
		r.state = ResultStateHasResults
		return nil

	case action.ToggleResultHeaders:
		r.hideHeaders = !r.hideHeaders
		return nil

	case action.UpdateSending:
		r.state = ResultStateIsSending
		return tick()

	case tickMsg:
		if r.state != ResultStateIsSending {
			r.sendAnimationFrame = 0
			return nil
		}
		r.sendAnimationFrame++
		if r.sendAnimationFrame >= len(sendAnimation) {
			r.sendAnimationFrame = 0
		}
		return tick()
	}

	return tea.Batch(r.toggleHeadersButton.Update(msg))
}

func tick() tea.Cmd {
	return tea.Tick(time.Second/24, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
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
