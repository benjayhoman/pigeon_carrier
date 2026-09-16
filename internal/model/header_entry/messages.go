package headerentry

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/pigeon_carrier/internal/action"
)

type HeaderEntryRemoveOnEnter struct {
	id uuid.UUID
}

func (h HeaderEntryRemoveOnEnter) Do() tea.Cmd {
	removeHeaderAction := action.NewRemoveHeader(h.id)
	return func() tea.Msg { return removeHeaderAction }
}
