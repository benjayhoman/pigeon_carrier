package program

import (
	"github.com/pigeon_carrier/internal/action"
	"github.com/pigeon_carrier/internal/app/model"

	tea "charm.land/bubbletea/v2"
)

func (m Program) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+q", "esc":
			return m, tea.Quit

		case "tab":
			return m.incrementAndSetFocus(1)

		case "shift+tab":
			return m.incrementAndSetFocus(-1)

		case "ctrl+s":
			m.resultState = ResultStateIsSending

			url := m.request.GetURL()
			method := m.request.GetMethod()
			headers := m.request.GetHeaders()
			body := m.request.GetBody()
			scriptPath := m.request.GetScriptPath()
			clientCertPath := m.request.GetClientCertificatePath()
			clientKeyPath := m.request.GetClientKeyPath()
			caCertPath := m.request.GetCaCertificatePath()

			return m, tea.Batch(action.NewDefaultMsg(action.UpdateSending{}),
				func() tea.Msg {
					return m.httpClient.CallHttp(m.log, model.HttpRequest{
						Url:                   url,
						Method:                method,
						Headers:               headers,
						Body:                  body,
						ScriptPath:            scriptPath,
						ClientCertificatePath: clientCertPath,
						ClientKeyPath:         clientKeyPath,
						CaCertificatePath:     caCertPath,
					})
				})
		}

	case action.UpdateFocus:
		return m.incrementAndSetFocus(msg.Increment)

	case action.UpdateResults:
		m.resultState = ResultStateHasResults
	}

	sendTabCmd := m.sendTab.Update(msg)
	loadTabCmd := m.loadTab.Update(msg)
	envTabCmd := m.envTab.Update(msg)

	requestCmd := m.request.Update(msg)
	resultsCmd := m.results.Update(msg)
	sendingCmd := m.sending.Update(msg)
	return m, tea.Batch(sendTabCmd, loadTabCmd, envTabCmd, requestCmd, resultsCmd, sendingCmd)
}
