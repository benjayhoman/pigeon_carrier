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
		case "ctrl+q", "esc":
			m.log.Debug().Msg("captured exit command")
			return m, tea.Quit

		case "tab":
			m.log.Debug().Msg("captured increment focus")
			return m.incrementAndSetFocus(1)

		case "shift+tab":
			m.log.Debug().Msg("captured decrement focus")
			return m.incrementAndSetFocus(-1)

		case "ctrl+s":
			m.log.Debug().Msg("captured send command")
			m.resultState = ResultStateIsSending

			url := m.urlInput.Value()
			headers := m.headers.GetHeaders()
			return m, tea.Batch(action.NewDefaultMsg(action.UpdateSending{}),
				func() tea.Msg {
					return httpclient.CallHttp(url, headers)
				})
		}

	case action.UpdateFocus:
		m.log.Debug().Msg("captured update focus action")
		return m.incrementAndSetFocus(msg.Increment)

	case action.UpdateResults:
		m.log.Debug().Msg("captured update results action")
		m.resultState = ResultStateHasResults
	}

	urlInputCmd := m.urlInput.Update(msg)
	methodCmd := m.method.Update(msg)
	headerCmd := m.headers.Update(msg)
	resultsCmd := m.results.Update(msg)
	sendingCmd := m.sending.Update(msg)
	return m, tea.Batch(urlInputCmd, methodCmd, headerCmd, resultsCmd, sendingCmd)
}
