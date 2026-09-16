package headers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pigeon_carrier/internal/action"
)

type AddNewHeaderOnEnter struct{}

func (a AddNewHeaderOnEnter) Do() tea.Cmd {
	addNewHeaderAction := action.NewAddHeader()
	return func() tea.Msg { return addNewHeaderAction }
}
