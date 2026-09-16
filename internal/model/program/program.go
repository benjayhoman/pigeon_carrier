package program

import (
	"github.com/pigeon_carrier/internal/model/headers"
	"github.com/pigeon_carrier/internal/model/method"
	"github.com/pigeon_carrier/internal/model/results"
	"github.com/pigeon_carrier/internal/model/urlinput"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Program struct {
	currentFocus int
	urlInput     *urlinput.UrlInput
	method       *method.Method
	headers      *headers.Headers
	results      *results.Results
}

func NewProgram() Program {
	method := method.NewMethod()
	urlInput := urlinput.NewUrlInput(method)
	headers := headers.NewHeaders()
	results := results.NewResults()
	program := Program{
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
