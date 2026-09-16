package headers

import (
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/add_remove_box"
	headerentry "github.com/pigeon_carrier/internal/model/header_entry"

	tea "charm.land/bubbletea/v2"
)

type Headers struct {
	headers      []headerentry.Entry
	addNewHeader *addremovebox.AddRemoveBox
}

func NewHeaders() *Headers {
	addNewHeader := addremovebox.NewAddRemoveBox("+", AddNewHeaderOnEnter{})
	return &Headers{
		headers:      make([]headerentry.Entry, 0),
		addNewHeader: &addNewHeader,
	}
}

func (h Headers) Init() tea.Cmd {
	return nil
}

func (h Headers) GetFocusables() []app.Focusable {
	focusables := make([]app.Focusable, 0, len(h.headers)*3)

	focusables = append(focusables, h.addNewHeader)
	for _, header := range h.headers {
		for _, focusable := range header.GetFocusables() {
			focusables = append(focusables, focusable)
		}
	}
	return focusables
}
