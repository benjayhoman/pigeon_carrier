package request

import (
	"fmt"
	"strings"
)

const (
	line = "─"
)

func (r Request) View() string {
	return fmt.Sprintf("Request: %s %s \n%s\n  %s\n  %s\n%s",
		r.saveButton.View(),
		r.envButton.View(),
		strings.Repeat(line, r.width),
		r.urlInput.View(),
		r.headers.View(),
		r.requestBody.View())
}
