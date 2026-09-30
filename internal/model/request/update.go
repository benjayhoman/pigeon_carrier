package request

import (
	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
)

func (r *Request) Update(msg tea.Msg) tea.Cmd {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.width = msg.Width

	case action.ToggleRequestBody:
		r.showRequestBody = !r.showRequestBody
		if r.showRequestBody {
			r.showBodyButton.SetSymbol("-")
		} else {
			r.showBodyButton.SetSymbol("+")
		}
	}

	saveButtonCmd := r.saveButton.Update(msg)
	envButtonCmd := r.envButton.Update(msg)
	urlInputCmd := r.urlInput.Update(msg)
	configCmd := r.config.Update(msg)
	headerCmd := r.headers.Update(msg)
	showBodyButtonCmd := r.showBodyButton.Update(msg)
	requestBodyCmd := r.requestBody.Update(msg)
	return tea.Batch(urlInputCmd, configCmd, headerCmd, requestBodyCmd, saveButtonCmd, envButtonCmd, showBodyButtonCmd)
}
