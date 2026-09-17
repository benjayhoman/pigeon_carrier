package headerentry

import (
	tea "charm.land/bubbletea/v2"
)

func (e *Entry) Update(msg tea.Msg) tea.Cmd {
	removeCmd := e.removeHeader.Update(msg)
	keyCmd := e.keyInput.Update(msg)
	valueCmd := e.valueInput.Update(msg)
	return tea.Batch(removeCmd, keyCmd, valueCmd)
}
