package program

import (
	"github.com/pigeon_carrier/internal/action"
	"github.com/pigeon_carrier/internal/model/headers"
	"github.com/pigeon_carrier/internal/model/method"
	"github.com/pigeon_carrier/internal/model/results"
	"github.com/pigeon_carrier/internal/model/sending"
	"github.com/pigeon_carrier/internal/model/urlinput"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

type ResultState string

const (
	ResultStateNone       ResultState = ""
	ResultStateHasResults ResultState = "hasResults"
	ResultStateIsSending  ResultState = "isSending"
)

type HttpClient interface {
	CallHttp(url string, headers map[string]string) action.UpdateResults
}

type Program struct {
	log          zerolog.Logger
	currentFocus int

	urlInput *urlinput.UrlInput
	method   *method.Method
	headers  *headers.Headers

	sending     *sending.Sending
	results     *results.Results
	resultState ResultState

	httpClient HttpClient
}

func NewProgram(logger zerolog.Logger, httpClient HttpClient) Program {
	programLogger := logger.With().Str("module", "program").Logger()

	method := method.NewMethod(programLogger)
	urlInput := urlinput.NewUrlInput(programLogger, method)
	headers := headers.NewHeaders(programLogger)
	sending := sending.NewSending(programLogger)
	results := results.NewResults(programLogger)

	program := Program{
		log:          programLogger,
		currentFocus: 0,
		method:       method,
		urlInput:     urlInput,
		headers:      headers,
		sending:      sending,
		results:      results,
		resultState:  ResultStateNone,
		httpClient:   httpClient,
	}
	program.setFocus() // set initial focus to the first focusable element

	program.log.Debug().Msg("Program initialized")
	return program
}

func (m Program) Init() tea.Cmd {
	return textinput.Blink
}
