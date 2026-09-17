package program

import (
	"github.com/pigeon_carrier/internal/action"
	"github.com/pigeon_carrier/internal/app/httpclient"

	tea "charm.land/bubbletea/v2"
)

func (m Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab":
			return m.incrementAndSetFocus(1)

		case "shift+tab":
			return m.incrementAndSetFocus(-1)

		case "ctrl+s":
			url := m.urlInput.Value()
			headers := m.headers.GetHeaders()
			return m, tea.Batch(action.NewDefaultMsg(action.UpdateSending{}),
				func() tea.Msg {
					return httpclient.CallHttp(url, headers)
				})
		}

	case action.UpdateFocus:
		return m.incrementAndSetFocus(msg.Increment)
	}

	urlInputCmd := m.urlInput.Update(msg)
	methodCmd := m.method.Update(msg)
	headerCmd := m.headers.Update(msg)
	resultsCmd := m.results.Update(msg)
	return m, tea.Batch(tea.ClearScreen, urlInputCmd, methodCmd, headerCmd, resultsCmd)
}
