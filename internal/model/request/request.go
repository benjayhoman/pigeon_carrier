package request

import (
	"github.com/pigeon_carrier/internal/app"
	"github.com/pigeon_carrier/internal/model/enterbox"
	"github.com/pigeon_carrier/internal/model/headers"
	"github.com/pigeon_carrier/internal/model/textfield"
	"github.com/pigeon_carrier/internal/model/urlinput"
	"github.com/rs/zerolog"

	tea "charm.land/bubbletea/v2"
)

type Request struct {
	saveButton *enterbox.EnterBox
	envButton  *enterbox.EnterBox

	urlInput *urlinput.UrlInput
	headers  *headers.Headers

	showBodyButton  *enterbox.EnterBox
	showRequestBody bool
	requestBody     *textfield.TextField

	width int
}

func NewRequest(logger zerolog.Logger) *Request {
	saveButton := enterbox.NewEnterBox(logger, "Save", nil)
	envButton := enterbox.NewEnterBox(logger, "Default Env", nil)
	showBodyButton := enterbox.NewEnterBox(logger, "+", ToggleRequestBodyOnEnter{})

	return &Request{
		urlInput:        urlinput.NewUrlInput(logger),
		headers:         headers.NewHeaders(logger),
		saveButton:      &saveButton,
		envButton:       &envButton,
		showBodyButton:  &showBodyButton,
		showRequestBody: false,
		requestBody:     textfield.NewTextField(),
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
	focusables := make([]app.Focusable, 0)
	focusables = append(focusables, r.saveButton)
	focusables = append(focusables, r.envButton)
	focusables = append(focusables, r.urlInput.GetFocusables()...)
	focusables = append(focusables, r.headers.GetFocusables()...)
	focusables = append(focusables, r.showBodyButton)
	focusables = append(focusables, r.requestBody)
	return focusables
}
