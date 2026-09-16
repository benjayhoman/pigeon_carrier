package results

import (
	"fmt"
)

func (r Results) View() string {
	if !r.hasResults {
		return ""
	}
	return fmt.Sprintf("%s %d\n%s\n",
		"Response: ",
		r.responseStatus,
		r.responseBody)
}
