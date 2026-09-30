package request

import (
	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
)

type ToggleRequestBodyOnEnter struct{}

func (t ToggleRequestBodyOnEnter) Do() tea.Cmd {
	return action.NewDefaultMsg(action.ToggleRequestBody{})
}
