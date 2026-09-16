package headerentry

import "fmt"

func (e Entry) View() string {
	return fmt.Sprintf("%s %s : %s", e.removeHeader.View(), e.keyInput.View(), e.valueInput.View())
}
