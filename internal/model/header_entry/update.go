package headerentry

import tea "github.com/charmbracelet/bubbletea"

func (e *Entry) Update(msg tea.Msg) tea.Cmd {
	removeCmd := e.removeHeader.Update(msg)
	keyCmd := e.keyInput.Update(msg)
	valueCmd := e.valueInput.Update(msg)
	return tea.Batch(removeCmd, keyCmd, valueCmd)
}
