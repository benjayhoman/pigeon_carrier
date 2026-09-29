package request

import tea "charm.land/bubbletea/v2"

func (r Request) Update(msg tea.Msg) tea.Cmd {
	urlInputCmd := r.urlInput.Update(msg)
	methodCmd := r.method.Update(msg)
	headerCmd := r.headers.Update(msg)
	requestBodyCmd := r.requestBody.Update(msg)
	return tea.Batch(urlInputCmd, methodCmd, headerCmd, requestBodyCmd)
}
