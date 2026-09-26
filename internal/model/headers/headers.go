package headers

import (
	"github.com/pigeon_carrier/internal/app"
	addremovebox "github.com/pigeon_carrier/internal/model/enterbox"
	headerentry "github.com/pigeon_carrier/internal/model/header_entry"
	"github.com/rs/zerolog"

	tea "charm.land/bubbletea/v2"
)

type Headers struct {
	log          zerolog.Logger
	headers      []headerentry.Entry
	addNewHeader *addremovebox.EnterBox
}

func NewHeaders(logger zerolog.Logger) *Headers {
	headersLogger := logger.With().Str("module", "headers").Logger()
	addNewHeader := addremovebox.NewEnterBox(headersLogger, "+", AddNewHeaderOnEnter{})
	return &Headers{
		log:          headersLogger,
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

func (h Headers) GetHeaders() map[string]string {
	headersMap := make(map[string]string)
	for _, header := range h.headers {
		key, value := header.GetKeyValue()
		headersMap[key] = value
	}
	return headersMap
}
