package sending

import "fmt"

func (s Sending) View() string {
	return fmt.Sprintf("%s %s \n\n", "", sendAnimation[s.frame])
}
