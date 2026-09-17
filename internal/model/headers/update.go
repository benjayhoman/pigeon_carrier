package headers

import (
	"slices"

	tea "charm.land/bubbletea/v2"
	"github.com/pigeon_carrier/internal/action"
	headerentry "github.com/pigeon_carrier/internal/model/header_entry"
)

func (h *Headers) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case action.AddHeader:
		header := headerentry.NewEntry()
		h.headers = append(h.headers, header)

		updateFocus := action.UpdateFocus{Increment: 0}
		return action.NewDefaultMsg(updateFocus)

	case action.RemoveHeader:
		h.headers = slices.DeleteFunc(h.headers, func(h headerentry.Entry) bool {
			return h.Id == msg.Id
		})

		updateFocus := action.UpdateFocus{Increment: 0}
		// Removing a header shrinks the inline view; clear the screen so the
		// renderer repaints from the top instead of leaving a stale first line.
		return tea.Batch(tea.ClearScreen, action.NewDefaultMsg(updateFocus))
	}

	addHeaderCmd := h.addNewHeader.Update(msg)

	var cmds = make([]tea.Cmd, 0, len(h.headers)*3)
	for _, header := range h.headers {
		cmds = append(cmds, header.Update(msg))
	}

	return tea.Batch(addHeaderCmd, tea.Batch(cmds...))
}
