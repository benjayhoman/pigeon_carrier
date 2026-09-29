package request

import "fmt"

func (r Request) View() string {
	return fmt.Sprintf("%s\n%s\n%s", r.urlInput.View(), r.headers.View(), r.requestBody.View())
}
