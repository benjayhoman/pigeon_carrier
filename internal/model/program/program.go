package program

import (
	"github.com/pigeon_carrier/internal/model/headers"
	"github.com/pigeon_carrier/internal/model/method"
	"github.com/pigeon_carrier/internal/model/results"
	"github.com/pigeon_carrier/internal/model/urlinput"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"github.com/rs/zerolog"
)

type Program struct {
	log          zerolog.Logger
	currentFocus int
	urlInput     *urlinput.UrlInput
	method       *method.Method
	headers      *headers.Headers
	results      *results.Results
}

func NewProgram(logger zerolog.Logger) Program {
	programLogger := logger.With().Str("module", "program").Logger()

	method := method.NewMethod(programLogger)
	urlInput := urlinput.NewUrlInput(programLogger, method)
	headers := headers.NewHeaders(programLogger)
	results := results.NewResults(programLogger)
	program := Program{
		log:          programLogger,
		currentFocus: 0,
		method:       method,
		urlInput:     urlInput,
		headers:      headers,
		results:      results,
	}
	program.setFocus() // set initial focus to the first focusable element
	return program
}

func (m Program) Init() tea.Cmd {
	return textinput.Blink
}
