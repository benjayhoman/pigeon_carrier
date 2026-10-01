package headers

import (
	"fmt"
)

func (h Headers) View() string {
	result := fmt.Sprintf("%s %s\n", "Headers:", h.addNewHeader.View())
	for _, header := range h.headers {
		result += fmt.Sprintf("    %s\n", header.View())
	}
	return result
}
