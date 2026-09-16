package headers

import (
	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
)

type AddNewHeaderOnEnter struct{}

func (a AddNewHeaderOnEnter) Do() tea.Cmd {
	addNewHeaderAction := action.NewAddHeader()
	return func() tea.Msg { return addNewHeaderAction }
}
