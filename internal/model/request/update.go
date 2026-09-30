package request

import tea "charm.land/bubbletea/v2"

func (r *Request) Update(msg tea.Msg) tea.Cmd {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.width = msg.Width
	}

	saveButtonCmd := r.saveButton.Update(msg)
	envButtonCmd := r.envButton.Update(msg)
	urlInputCmd := r.urlInput.Update(msg)
	headerCmd := r.headers.Update(msg)
	requestBodyCmd := r.requestBody.Update(msg)
	return tea.Batch(urlInputCmd, headerCmd, requestBodyCmd, saveButtonCmd, envButtonCmd)
}
