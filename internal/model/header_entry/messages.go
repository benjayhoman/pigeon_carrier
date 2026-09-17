package headerentry

import (
	tea "charm.land/bubbletea/v2"
	"github.com/google/uuid"
	"github.com/pigeon_carrier/internal/action"
)

type HeaderEntryRemoveOnEnter struct {
	id uuid.UUID
}

func (h HeaderEntryRemoveOnEnter) Do() tea.Cmd {
	removeHeaderAction := action.NewRemoveHeader(h.id)
	return action.NewDefaultMsg(removeHeaderAction)
}
