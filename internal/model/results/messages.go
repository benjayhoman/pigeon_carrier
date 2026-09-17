package results

import (
	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
)

type ToggleResultHeadersOnEnter struct{}

func (t ToggleResultHeadersOnEnter) Do() tea.Cmd {
	return action.NewDefaultMsg(action.ToggleResultHeaders{})
}
