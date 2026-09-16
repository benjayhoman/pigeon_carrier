package headers

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pigeon_carrier/internal/action"
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/add_remove_box"
	headerentry "github.com/pigeon_carrier/internal/model/header_entry"
)

type AddNewHeaderOnEnter struct{}

func (a AddNewHeaderOnEnter) Do() tea.Cmd {
	addNewHeaderAction := action.NewAddHeader()
	return func() tea.Msg { return addNewHeaderAction }
}

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
	boxes := make([]app.Focusable, 0, len(h.headers)*3)
	for _, header := range h.headers {
		for _, focusable := range header.GetFocusables() {
			boxes = append(boxes, focusable)
		}
	}
	boxes = append(boxes, h.addNewHeader)
	return boxes
}
