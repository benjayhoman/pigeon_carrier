package request

import (
	"fmt"
	"strings"
)

const (
	line = "─"
)

func (r Request) View() string {
	return fmt.Sprintf("Request: %s %s \n%s\n%s\n  %s\n  %s  Body: %s\n%s\n",
		r.saveButton.View(),
		r.envButton.View(),
		strings.Repeat(line, r.width),
		r.urlInput.View(),
		r.config.View(),
		r.headers.View(),
		r.showBodyButton.View(),
		r.resolveRequestBodyView())
}

func (r Request) resolveRequestBodyView() string {
	if r.showRequestBody {
		return r.requestBody.View()
	}
	return ""
}
