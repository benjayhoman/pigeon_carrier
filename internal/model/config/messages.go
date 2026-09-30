package config

import (
	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
)

type ToggleConfigOnEnter struct{}

func (t ToggleConfigOnEnter) Do() tea.Cmd {
	return action.NewDefaultMsg(action.ToggleConfig{})
}
