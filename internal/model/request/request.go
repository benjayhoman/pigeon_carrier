package request

import (
	"github.com/pigeon_carrier/internal/app"
	"github.com/pigeon_carrier/internal/model/headers"
	"github.com/pigeon_carrier/internal/model/method"
	"github.com/pigeon_carrier/internal/model/textfield"
	"github.com/pigeon_carrier/internal/model/urlinput"
	"github.com/rs/zerolog"

	tea "charm.land/bubbletea/v2"
)

type Request struct {
	method      *method.Method
	urlInput    *urlinput.UrlInput
	headers     *headers.Headers
	requestBody *textfield.TextField
}

func NewRequest(logger zerolog.Logger) *Request {
	method := method.NewMethod(logger)

	return &Request{
		method:      method,
		urlInput:    urlinput.NewUrlInput(logger, method),
		headers:     headers.NewHeaders(logger),
		requestBody: textfield.NewTextField(),
	}
}

func (r Request) Init() tea.Cmd {
	return nil
}

func (r Request) GetURL() string {
	return r.urlInput.Value()
}

func (r Request) GetHeaders() map[string]string {
	return r.headers.GetHeaders()
}

func (r Request) GetFocusables() []app.Focusable {
	focusables := make([]app.Focusable, 0, len(r.headers.GetFocusables())+3)
	focusables = append(focusables, r.method, r.urlInput)
	focusables = append(focusables, r.headers.GetFocusables()...)
	focusables = append(focusables, r.requestBody)
	return focusables
}